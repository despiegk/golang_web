package dns

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/miekg/dns"
)

var db *bolt.DB

func DNSDBStart() {
	var err error
	db, err = bolt.Open("dns.db", 0600, nil)
	if err != nil {
		log.Fatalf("Failed to open Bolt DB: %v", err)
	}
}

func RegisterARecord(name, ip string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("dns"))
		if err != nil {
			return err
		}
		log.Printf("Register IP addr %s ", name)
		return b.Put([]byte(name), []byte(ip))
	})
}

func DnsStart() {

	mux := dns.NewServeMux()
	mux.HandleFunc(".", handleDNSRequest)

	server := &dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: mux,
	}

	log.Println("Starting DNS server on :53")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %s\n", err.Error())
	}

}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	switch r.Opcode {
	case dns.OpcodeQuery:
		for _, q := range r.Question {
			switch q.Qtype {
			case dns.TypeA:
				log.Printf("Handling A record request for %s", q.Name)
				ip, err := getARecord(q.Name)
				if err == nil && ip != "" {
					log.Printf("Found A record for %s: %s", q.Name, ip)
					rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
					if err == nil {
						msg.Answer = append(msg.Answer, rr)
					}
				} else {
					log.Printf("A record for %s not found, forwarding request to 8.8.8.8", q.Name)
					// Forward the request to 8.8.8.8 if not known in Bolt
					c := new(dns.Client)
					m := new(dns.Msg)
					m.SetQuestion(q.Name, q.Qtype)
					resp, _, err := c.Exchange(m, "8.8.8.8:53")
					if err == nil && resp != nil && len(resp.Answer) > 0 {
						msg.Answer = resp.Answer
					} else {
						log.Printf("Error forwarding DNS request: %v", err)
					}
				}
			}
		}
	}

	log.Printf("Responding to DNS request with %d answers", len(msg.Answer))
	w.WriteMsg(&msg)
}

func getARecord(name string) (string, error) {
	var ip string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dns"))
		if b != nil {
			if name[len(name)-1] == '.' {
				name = name[:len(name)-1]
			}
			value := b.Get([]byte(name))
			if value != nil {
				ip = string(value)
				log.Printf("DNS get: %s %s", name, ip)
			} else {
				return fmt.Errorf("a record for %s not found", name)
			}
		}
		return nil
	})
	return ip, err
}

func ListARecords() ([]string, error) {
	var records []string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("dns"))
		if b != nil {
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				records = append(records, fmt.Sprintf("%s A %s", k, v))
			}
		}
		return nil
	})
	return records, err
}
