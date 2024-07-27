"use client";

import React, { createContext, useState, useEffect, ReactNode } from "react";
import { Book } from "@/types/Book";

import {
  getBooks,
  postBook,
  editBook as editBookOnServer,
  deleteBook as deleteBookOnServer,
} from "@/services/bookService";

interface BookContextType {
  books: Book[];
  addBook: (book: Book) => void;
  editBook: (book: Book) => void;
  deleteBook: (id: number) => void;
  error: string | null;
}

const BookContext = createContext<BookContextType | undefined>(undefined);

const BookProvider = ({ children }: { children: ReactNode }) => {
  const [books, setBooks] = useState<Book[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const data = await getBooks();
        setBooks(data);
      } catch (error) {
        setError("Failed to fetch books. Please try again later.");
        console.error(error);
      }
    };

    fetchBooks();
  }, []);

  const addBook = async (book: Book) => {
    try {
      const newBook = await postBook(book);
      setBooks([...books, newBook]);
      setError(null);
    } catch (error) {
      setError("Failed to add book. Please try again.");
      console.error(error);
    }
  };

  const editBook = async (updatedBook: Book) => {
    try {
      const updatedBookFromServer = await editBookOnServer(updatedBook);

      setBooks(
        books.map((book) =>
          book.id === updatedBookFromServer.id ? updatedBookFromServer : book
        )
      );
      setError(null);
    } catch (error) {
      setError("Failed to update book. Please try again.");
      console.error(error);
    }
  };

  const deleteBook = async (id: number) => {
    try {
      await deleteBookOnServer(id);
      setBooks(books.filter((book) => book.id !== id));
      setError(null);
    } catch (error) {
      setError("Failed to delete book. Please try again.");
      console.error(error);
    }
  };

  return (
    <BookContext.Provider
      value={{ books, addBook, editBook, deleteBook, error }}
    >
      {children}
      {error && (
        <div className="text-red-700 bg-red-100 border border-red-400 p-4 rounded mt-4 shadow-md">
          <p>{error}</p>
        </div>
      )}
    </BookContext.Provider>
  );
};

export { BookContext, BookProvider };
