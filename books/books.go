package books

type Book struct {
	ID            int    `json:"courseId"`
	Title         string `json:"bookName"`
	Author        string `json:"writer"`
	YearPublished int    `json:"published year"`
}

var books = []Book{
	{
		ID:            1,
		Title:         "Introduction to Linux",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R Tolkein",
		YearPublished: 1937,
	},
	{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles",
		YearPublished: 2009,
	},
	{
		ID:            4,
		Title:         "Harry Potter and the Philosophers Stone",
		Author:        "J.K. Rowling",
		YearPublished: 1997,
	},
	{
		ID:            5,
		Title:         "Les Miserables",
		Author:        "Victor Hugo",
		YearPublished: 1862,
	},
}