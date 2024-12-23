package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type BookItem struct {
	Code        int    `json:"code"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Img         string `json:"img"`
	IsFavorite  bool   `json:"is_favorite"`
}

var booksStore = []BookItem{
	{Code: 1343, Title: "1984", Author: "Джордж Оруэлл", Price: 800, Img: "1.jpg", Description: "«1984» (англ. Nineteen Eighty-Four, «тысяча девятьсот восемьдесят четвёртый») — роман-антиутопия Джорджа Оруэлла, изданный в 1949 году. Последнее и, по господствующему мнению, самое главное произведение писателя", IsFavorite: false},
	{Code: 22342, Title: "Убить пересмешника", Author: "Харпер Ли", Price: 700, Img: "2.jpg", Description: "«Уби́ть пересме́шника» (англ. To Kill a Mockingbird) — роман-бестселлер американской писательницы Харпер Ли, опубликованный в 1960 году, за который в 1961 году она получила Пулитцеровскую премию. Её успех стал вехой в борьбе за права чернокожих.", IsFavorite: true},
	{Code: 3222, Title: "Гордость и предубеждение", Author: "Джейн Остин", Price: 600, Img: "3.jpg", Description: "Роман начинается с беседы мистера и миссис Беннет о приезде молодого мистера Бингли в Незерфилд-парк. Жена уговаривает мужа навестить соседа и свести с ним более тесное знакомство. Она надеется, что мистеру Бингли непременно понравится одна из их пяти дочерей. Мистер Беннет наносит визит молодому человеку, и тот через какое-то время наносит ответный.", IsFavorite: false},
	{Code: 496754, Title: "Моби Дик", Author: "Herman Melville", Price: 750, Img: "4.jpg", Description: "Повествование ведётся от имени американского моряка Измаила, ушедшего в рейс на китобойном судне «Пекод», капитан которого, Ахав (отсылка к библейскому Ахаву), одержим идеей мести гигантскому белому киту, убийце китобоев, известному как Моби Дик (в предыдущем плавании по вине кита Ахав потерял ногу, и с тех пор капитан использует протез).", IsFavorite: false},
	{Code: 5453, Title: "Великий Гэтсби", Author: "Фрэнсис Скотт Фицджеральд", Price: 650, Img: "2.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 63578, Title: "Преступление и наказание", Author: "Фёдор Достоевский", Price: 900, Img: "1.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 732234, Title: "Анна Каренина", Author: "Лев Толстой", Price: 850, Img: "4.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 85843, Title: "Сияние", Author: "Стивен Кинг", Price: 800, Img: "3.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 94622, Title: "Три товарища", Author: "Эрих Мария Ремарк", Price: 700, Img: "1.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 103468, Title: "451 градус по Фаренгейту", Author: "Рэй Брэдбери", Price: 600, Img: "4.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 119275, Title: "Братья Карамазовы", Author: "Фёдор Достоевский", Price: 950, Img: "2.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 125006, Title: "Маленький принц", Author: "Антуан де Сент-Экзюпери", Price: 500, Img: "3.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 13227, Title: "Старик и море", Author: "Эрнест Хемингуэй", Price: 700, Img: "1.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 14, Title: "Война и мир", Author: "Лев Толстой", Price: 1200, Img: "2.jpg", IsFavorite: false},
	{Code: 152342, Title: "Дон Кихот", Author: "Мигель де Сервантес", Price: 800, Img: "4.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 167732, Title: "Собачье сердце", Author: "Михаил Булгаков", Price: 600, Img: "3.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 172345, Title: "Над пропастью во ржи", Author: "Джером Д. Сэлинджер", Price: 700, Img: "1.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 185674, Title: "Цветы для Элджернона", Author: "Дэниел Киз", Price: 650, Img: "4.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
	{Code: 194178, Title: "Тень ветра", Author: "Карлос Руис Сафон", Price: 800, Img: "2.jpg", Description: "Итоговое произведение литературы американского романтизма. Длинный роман с многочисленными лирическими отступлениями, проникнутый библейской образностью и многослойным символизмом, не был понят и принят современниками", IsFavorite: false},
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/books", getBooks)
	router.POST("/books/setfavorite/:code", setFavorite)
	router.GET("/books/info/:code", getBook)
	router.POST("/books", addBook)
	router.DELETE("/books/:code", deleteBook)

	router.Run("localhost:3000")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, booksStore)
}

func getBook(c *gin.Context) {
	codeStr := c.Param("code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid code"})
		return
	}
	for _, book := range booksStore {
		if book.Code == code {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var newBook BookItem
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booksStore = append(booksStore, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
	codeStr := c.Param("code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid code"})
		return
	}
	for index, book := range booksStore {
		if book.Code == code {
			booksStore = append(booksStore[:index], booksStore[index+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func setFavorite(c *gin.Context) {
	codeStr := c.Param("code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid code"})
		return
	}
	for index, book := range booksStore {
		if book.Code == code {
			booksStore[index].IsFavorite = !booksStore[index].IsFavorite
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
