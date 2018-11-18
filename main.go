package main

import (
	"github.com/nats-io/go-nats-streaming"
	"github.com/riverrco/Notification-Service/pkg/queue"
	"github.com/urfave/cli"
	"os"
	"log"
)

func main(){
	app := cli.NewApp()
	app.Name = "notification-service"
	app.Usage = "NGripp Notification Service"
	app.Action = start

	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "NATS_HOST",
			Name:   "NATS_HOST",
			Value:  "nats://localhost:4223",
			Usage:  "Nats host",
		},
		cli.StringFlag{
			EnvVar: "NATS_CLUSTER_ID",
			Name:   "NATS_CLUSTER_ID",
			Value:  "test-cluster",
			Usage:  "Nats Cluster ID",
		},
		cli.StringFlag{
			EnvVar: "NATS_CLIENT_ID",
			Name:   "NATS_CLIENT_ID",
			Value:  "queue-notifications",
			Usage:  "Nats Client ID",
		},
		cli.StringFlag{
			EnvVar: "MAILGUN_DOMAIN",
			Name:   "MAILGUN_DOMAIN",
			Value:  "sandbox88fdc674a0844d7d959008b82d0b7bd4.mailgun.org",
			Usage:  "Mailgun Domain",
		},
		cli.StringFlag{
			EnvVar: "MAILGUN_PRIVATE_KEY",
			Name:   "MAILGUN_PRIVATE_KEY",
			Value:  "",
			Usage:  "Mailgun private key",
		},
		cli.StringFlag{
			EnvVar: "MAILGUN_PUBLIC_KEY",
			Name:   "MAILGUN_PUBLIC_KEY",
			Value:  "",
			Usage:  "Mailgun public key",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context){
	conn, err := stan.Connect(
		c.String("NATS_CLUSTER_ID"),
		c.String("NATS_CLIENT_ID"),
		stan.NatsURL(c.String("NATS_HOST")),
		)

	if err != nil {
		log.Printf("Error on Nats %v", err)
	}
	cl := queue.NewNats(conn)
	err = cl.Sub("notifiy", "notification-1",  listener)

	if err != nil {
		log.Printf("Error on Sub %v", err)
	}


}