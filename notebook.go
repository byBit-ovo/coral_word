package main

import(
	"fmt"
	"github.com/google/uuid"
)


func createNoteBook(session string, bookName string) (err error){
	uid,ok := userSession[session]
	if ok == false{
		return fmt.Errorf("user isn't logged in !")
	}
	tx, err := db.Begin()
	if err != nil{
		return err
	}
	defer func(){
		if err != nil{
			_ = tx.Rollback()
		}
	}()
	book_id := uuid.New().String()
	_, err = tx.Exec("insert into note_book (book_id, book_name, user_id) values (?,?,?)",book_id, bookName,uid)
	if err != nil{
		return fmt.Errorf("failed to insert notebook: %w", err)
	}
	return tx.Commit()
}

func AddWordToNotebook(session, word, noteBookName string) (err error){
	uid,ok := userSession[session]
	if ok == false{
		return fmt.Errorf("user isn't logged in !")
	}
	tx, err := db.Begin()
	if err != nil{
		return err
	}
	defer func(){
		if err != nil{
			_ = tx.Rollback()
		}
	}()
	row := tx.QueryRow("select id from vocabulary where word = ?", word)
	var wordId int64
	if err:=row.Scan(&wordId); err != nil{
		return err
	}
	row = tx.QueryRow("select book_id from note_book where book_name=? and user_id=?", noteBookName, uid)
	var bookId int64
	if err:=row.Scan(&bookId); err != nil{
		return err
	}
	_, err = tx.Exec("insert into note_word (book_id, word_id) values (?,?)", bookId, wordId)
	if err != nil{
		return err
	}

	return tx.Commit()
}