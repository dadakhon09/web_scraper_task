package handlers

import (
	"github.com/dadakhon09/web_scraper_task/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"net/http"
	"sync"
)

// @Router /task [post]
// @Summary RESTful API endpoint
// @Description Takes an integer that represents the number of threads
// @Tags task
// @Accept  json
// @Produce  json
// @Param number body models.Request true "number of threads/goroutines"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) TaskHandler(c *gin.Context) {
	var (
		req      models.Request
		resp     models.Response
		wg       sync.WaitGroup
		websites = []string{"https://www.result.si/projekti/", "https://www.result.si/o-nas/", "https://www.result.si/kariera/", "https://www.result.si/blog/"}
	)

	err := c.ShouldBind(&req)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	switch req.Number {
	case 1:
		wg.Add(1)

		go func() {
			result, errCounter, err := AllConsecutiveWebScraper(websites, &wg)
			if h.handleError(c, err, "Error on web scraping all pages consecutively") {
				return
			}
			resp = models.Response{
				NumOfSuccessCalls: int32(len(websites)) - errCounter,
				NumOfFailedCalls:  errCounter,
				Titles:            result,
			}
		}()

	case 2:
		wg.Add(2)

		go func() {
			title, errCounter, err := WebScraper(websites[0], &wg, false)
			if h.handleError(c, err, "Error on web scraping the first page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[0],
				Title: title,
			})

			title, errCounter2, err := WebScraper(websites[1], &wg, true)
			if h.handleError(c, err, "Error on web scraping the second page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[1],
				Title: title,
			})
			resp.NumOfFailedCalls = errCounter + errCounter2
			resp.NumOfSuccessCalls = int32(len(websites)) - (errCounter + errCounter2)
		}()

		go func() {
			title, errCounter, err := WebScraper(websites[2], &wg, false)
			if h.handleError(c, err, "Error on web scraping the third page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[2],
				Title: title,
			})

			title, errCounter2, err := WebScraper(websites[3], &wg, true)
			if h.handleError(c, err, "Error on web scraping the fourth page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[3],
				Title: title,
			})
			resp.NumOfFailedCalls = resp.NumOfFailedCalls + errCounter + errCounter2
			resp.NumOfSuccessCalls = resp.NumOfSuccessCalls + int32(len(websites)) - (errCounter + errCounter2)

		}()

	case 3:
		wg.Add(3)

		go func() {
			title, errCounter, err := WebScraper(websites[0], &wg, false)
			if h.handleError(c, err, "Error on web scraping the first page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[0],
				Title: title,
			})

			title, errCounter2, err := WebScraper(websites[1], &wg, true)
			if h.handleError(c, err, "Error on web scraping the second page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[1],
				Title: title,
			})
			resp.NumOfFailedCalls = errCounter + errCounter2
			resp.NumOfSuccessCalls = int32(len(websites)) - (errCounter + errCounter2)

		}()

		go func() {
			title, errCounter, err := WebScraper(websites[2], &wg, true)
			if h.handleError(c, err, "Error on web scraping the third page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[2],
				Title: title,
			})
			resp.NumOfFailedCalls = resp.NumOfFailedCalls + errCounter
			resp.NumOfSuccessCalls = resp.NumOfSuccessCalls + int32(len(websites)) - errCounter
		}()

		go func() {
			title, errCounter, err := WebScraper(websites[3], &wg, true)
			if h.handleError(c, err, "Error on web scraping the fourth page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[3],
				Title: title,
			})

			resp.NumOfFailedCalls = resp.NumOfFailedCalls + errCounter
			resp.NumOfSuccessCalls = resp.NumOfSuccessCalls + int32(len(websites)) - errCounter
		}()

	case 4:
		wg.Add(4)

		go func() {

			title, errCounter, err := WebScraper(websites[0], &wg, true)
			if h.handleError(c, err, "Error on web scraping the first page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[0],
				Title: title,
			})

			resp.NumOfFailedCalls = errCounter
			resp.NumOfSuccessCalls = int32(len(websites)) - errCounter

		}()

		go func() {

			title, errCounter, err := WebScraper(websites[1], &wg, true)
			if h.handleError(c, err, "Error on web scraping the second page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[1],
				Title: title,
			})

			resp.NumOfFailedCalls = resp.NumOfFailedCalls + errCounter
			resp.NumOfSuccessCalls = resp.NumOfSuccessCalls + int32(len(websites)) - errCounter
		}()

		go func() {

			title, errCounter, err := WebScraper(websites[2], &wg, true)
			if h.handleError(c, err, "Error on web scraping the third page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[2],
				Title: title,
			})
			resp.NumOfFailedCalls = resp.NumOfFailedCalls + errCounter
			resp.NumOfSuccessCalls = resp.NumOfSuccessCalls + int32(len(websites)) - errCounter
		}()

		go func() {

			title, errCounter, err := WebScraper(websites[3], &wg, true)
			if h.handleError(c, err, "Error on web scraping the fourth page") {
				return
			}
			resp.Titles = append(resp.Titles, models.Obj{
				Link:  websites[3],
				Title: title,
			})

			resp.NumOfFailedCalls = resp.NumOfFailedCalls + errCounter
			resp.NumOfSuccessCalls = resp.NumOfSuccessCalls + int32(len(websites)) - errCounter
		}()
	default:
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorBadRequest,
			Message: "the input should be in range (1,4)",
		})
		return
	}
	wg.Wait()

	c.JSON(http.StatusOK, resp)
}

func WebScraper(link string, wg *sync.WaitGroup, isDone bool) (title string, errCounter int32, err error) {
	var titles []string

	errCounter = 0
	c := colly.NewCollector()

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		titles = append(titles, e.Text)
	})

	err = c.Visit(link)
	if err != nil {
		errCounter++
	}
	if isDone {
		wg.Done()
	}
	return titles[0], 0, nil
}

func AllConsecutiveWebScraper(links []string, wg *sync.WaitGroup) (resp []models.Obj, errCounter int32, err error) {
	var (
		titles []string
	)
	resp = []models.Obj{}
	errCounter = 0
	c := colly.NewCollector()

	for i := 0; i < 4; i++ {
		titles = []string{}
		c.OnHTML("h2", func(e *colly.HTMLElement) {
			titles = append(titles, e.Text)
		})

		err = c.Visit(links[i])
		if err != nil {
			errCounter++
		}

		resp = append(resp, models.Obj{
			Link:  links[i],
			Title: titles[0],
		})
	}
	wg.Done()

	return
}
