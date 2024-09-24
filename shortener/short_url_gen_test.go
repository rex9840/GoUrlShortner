
package shortner 

import ( 
        "testing"
        "github.com/stretchr/testify/assert"
)


const userId = "eqweqwe-213-qwe" 

func TestShorrtLinkGen(T *testing.T){ 
        initialLink1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html" 
        userId := "e0dba740-fc4b-4977-872c-d360239e6b10"
        shortLink := GenerateShortURLLink(initialLink1, userId) 

        assert.Equal(T,shortLink,"9Zatkhpi")

}

