package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/cakirmuha/auction-bid-tracker/model"
)

const (
	url = "http://localhost:8181/api/v1"
)

func TestConcurrentBids(t *testing.T) {
	var tdata = []struct {
		inputItem      uint32
		expectedAmount uint64
	}{
		{1, 100},
		{2, 200},
		{3, 300},
		{4, 400},
	}
	for _, tc := range tdata {
		CreateBidsForAllMockUsers(tc.inputItem)
		outBidAmount := GetCurrentWinningBidAmountByItem(tc.inputItem)

		if outBidAmount == nil {
			t.Errorf("GetCurrentWinningBidAmountByItem(%v): Err", tc.inputItem)
			continue
		}
		if *outBidAmount != tc.expectedAmount {
			t.Errorf("GetCurrentWinningBidAmountByItem(%v): Got %v want %v", tc.inputItem, *outBidAmount, tc.expectedAmount)
		}
	}
}

func CreateBidsForAllMockUsers(itemID uint32) {
	bids := make(chan *model.Bid)
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			for b := range bids {
				CreateBid(b)
			}
			wg.Done()
		}()
	}
	for j := 1; j <= 100; j++ {
		b := &model.Bid{
			UserID: uint32(j),
			ItemID: itemID,
			Amount: uint64(j) * uint64(itemID),
		}
		bids <- b
	}
	close(bids)
	wg.Wait()
}

func GetCurrentWinningBidAmountByItem(itemID uint32) *uint64 {
	client := http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/item/%v/bids/winning", url, itemID), nil)
	if err != nil {
		log.Infof("request - current winning bid for item(%v) couldn't be created", itemID)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Infof("response - current winning bid for item(%v) couldn't be created", itemID)
		return nil
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Infof("read body - current winning bid for item(%v)", itemID)
		return nil
	}

	response := struct {
		Error model.ApiResponseError `json:"error"`
		Data  model.Bid              `json:"data"`
	}{}

	if err := json.Unmarshal(b, &response); err != nil {
		log.Infof("unmarshal bid - current winning bid for item(%v)", itemID)
		return nil
	}

	return &response.Data.Amount
}

func CreateBid(bid *model.Bid) {
	b, _ := json.Marshal(bid)
	payload := strings.NewReader(string(b))

	client := http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/user/bid", url), payload)
	if err != nil {
		log.Infof("request bid(user(%v)-item(%v)-amount(%v)) couldn't be created", bid.UserID, bid.ItemID, bid.Amount)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Infof("response bid(user(%v)-item(%v)-amount(%v)) couldn't be created", bid.UserID, bid.ItemID, bid.Amount)
		return
	}
	defer res.Body.Close()
}
