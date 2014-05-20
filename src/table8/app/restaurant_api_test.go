package app

import (
	"fmt"
	"net/http"
	"testing"
)

func TestEchosContent (t *testing.T) {
	{
	    req, err := http.NewRequest("GET", "/", nil)
	    reply := GeneralReply{}
	    System.Test(req, &struct{testing=1}, &reply)
	}
}
