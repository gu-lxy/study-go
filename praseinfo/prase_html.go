package praseinfo

import (
	"NewMovieSpider/entity"
	"regexp"
	"strings"
)


func ParseHtmlPage(pageString string)[]entity.Movie{

	idExp :=regexp.MustCompile(`<a class="nbg" href="https://movie.douban.com/subject/(.*?)/"`)

	idSlice := idExp.FindAllStringSubmatch(pageString,-1)

	nameExp := regexp.MustCompile(`title="(.*?)"`)
	nameSlice := nameExp.FindAllStringSubmatch(pageString,-1)

	descExp := regexp.MustCompile(`<p class="pl">(.*?)</p>`)
	descSlice := descExp.FindAllStringSubmatch(pageString,-1)


	rateNumExp := regexp.MustCompile(`<span class="rating_nums">(.*?)</span>`)
	rateNumSlice := rateNumExp.FindAllStringSubmatch(pageString,-1)

	voteNumExp := regexp.MustCompile(`<span class="pl">(.*?)</span>`)
	voteSlice := voteNumExp.FindAllStringSubmatch(pageString,-1)
	//处理数据，对采集的但是不符合要求的数据进行特殊的处理，使之符合要求
	for _,value := range voteSlice {
		//(30564人评价) ： value[1]
		value[1] = strings.ReplaceAll(value[1],"(","")
		value[1] = strings.ReplaceAll(value[1],"人评价)","")
	}


	coverUrlExp := regexp.MustCompile(`<img src="(.*?)" width="75" alt`)
	coverUrlSlice := coverUrlExp.FindAllStringSubmatch(pageString,-1)

	movies := make([]entity.Movie,0)
	for i := 0; i < len(idSlice); i++{

		movie := entity.Movie{
			Id:       idSlice[i][1],
			Name:     nameSlice[i][1],
			Desc:     descSlice[i][1],
			RateNum:  rateNumSlice[i][1],
			VoteNum:  voteSlice[i][1],
			CoverUrl: coverUrlSlice[i][1],
		}

		movies = append(movies,movie)
	}
	return movies
}
