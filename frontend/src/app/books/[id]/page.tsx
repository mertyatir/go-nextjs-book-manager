"use client";

import { useContext } from "react";
import { BookContext } from "../../../context/BookContext";

const BookDetail = ({ params }: { params: { id: string } }) => {
  const id = params.id;

  const context = useContext(BookContext);

  if (!context) {
    return <div>Context not available</div>;
  }

  const { books } = context;

  if (!id) return <div>Invalid book ID</div>;

  const bookId = parseInt(id as string);
  const book = books.find((book) => book.id === bookId);

  if (!book) return <div>Book not found</div>;

  return (
    <div className="max-w-2xl mx-auto p-6 bg-white shadow-md rounded-lg">
      <h1 className="text-3xl font-bold mb-4">{book.title}</h1>
      <p className="text-lg mb-2">
        <strong>Author:</strong> {book.author}
      </p>
      <p className="text-lg mb-2">
        <strong>Year:</strong> {book.year}
      </p>
      {book.genre && (
        <p className="text-lg mb-2">
          <strong>Genre:</strong> {book.genre}
        </p>
      )}
      {book.isbn && (
        <p className="text-lg mb-2">
          <strong>ISBN:</strong> {book.isbn}
        </p>
      )}
      {book.publisher && (
        <p className="text-lg mb-2">
          <strong>Publisher:</strong> {book.publisher}
        </p>
      )}
      {book.description && (
        <p className="text-lg mb-2">
          <strong>Description:</strong> {book.description}
        </p>
      )}
    </div>
  );
};

export default BookDetail;
