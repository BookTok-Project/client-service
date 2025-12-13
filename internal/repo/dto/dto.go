package dto

import "client-service/internal/db/pg"

type GetBooksLikeCountsArgs struct {
	BookIDs []int64
}

type GetBooksLikeCountsResult struct {
	Items []GetBooksLikeCountsResultItem
}

func ConvertToGetBooksLikeCountsResult(rows []pg.GetBooksLikeCountsRow) GetBooksLikeCountsResult {
	res := GetBooksLikeCountsResult{
		Items: make([]GetBooksLikeCountsResultItem, 0, len(rows)),
	}

	for _, row := range rows {
		res.Items = append(res.Items, ConvertToGetBooksLikeCountsResultItem(row))
	}

	return res
}

type GetBooksLikeCountsResultItem struct {
	BookID int64
	Count  int64
}

func ConvertToGetBooksLikeCountsResultItem(row pg.GetBooksLikeCountsRow) GetBooksLikeCountsResultItem {
	item := GetBooksLikeCountsResultItem{
		BookID: row.BookID,
		Count:  row.Count,
	}

	return item
}
