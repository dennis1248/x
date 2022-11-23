//go:build ignore

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"within.website/x/web/mastodon"
)

func promptInput(prompt string) (string, error) {
	fmt.Fprint(os.Stderr, prompt)
	fmt.Fprint(os.Stderr, ": ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input, nil
}

func main() {
	instance, err := promptInput("URL of the mastodon server (incl https://)")
	if err != nil {
		log.Fatal(err)
	}

	token, err := promptInput("Mastodon token")
	if err != nil {
		log.Fatal(err)
	}

	cli, err := mastodon.Authenticated("Xe/x test", "https://within.website/.x.botinfo", instance, token)
	if err != nil {
		log.Fatal(err)
	}

	if err := cli.VerifyCredentials(context.Background()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("your token works!")

	fname, err := promptInput("path to a jpeg")
	if err != nil {
		log.Fatal(err)
	}

	fin, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	descripiton, err := promptInput("description of image")
	if err != nil {
		log.Fatal(err)
	}

	att, err := cli.UploadMedia(context.Background(), fin, fname, descripiton, "")
	if err != nil {
		log.Fatal(err)
	}

	statusText, err := promptInput("toot")
	if err != nil {
		log.Fatal(err)
	}

	st, err := cli.CreateStatus(context.Background(), mastodon.CreateStatusParams{
		Status:   statusText,
		MediaIDs: []string{att.ID},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(st.URL)
}
