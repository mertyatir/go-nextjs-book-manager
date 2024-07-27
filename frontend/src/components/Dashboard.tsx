"use client";

import React, { useContext, useState } from "react";
import { BookContext } from "../context/BookContext";
import BookForm from "@/components/BookForm";
import Modal from "@/components/Modal";
import Link from "next/link";

import { Book } from "@/types/Book";

const Dashboard = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [currentBook, setCurrentBook] = useState<Book | undefined>(undefined);

  const context = useContext(BookContext);

  if (!context) {
    return <div>Context not available</div>;
  }

  const { books, deleteBook } = context;

  const handleEdit = (book: Book) => {
    setCurrentBook(book);
    setIsModalOpen(true);
  };

  return (
    <div className="container mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6">Book Dashboard</h1>
      <button
        onClick={() => {
          setCurrentBook(undefined);
          setIsModalOpen(true);
        }}
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mb-6"
      >
        Add Book
      </button>
      <ul className="space-y-4">
        {books.map((book) => (
          <li
            key={book.id}
            className="p-4 bg-white shadow-md rounded-lg flex justify-between items-center"
          >
            <div>
              <span className="font-semibold">{book.title}</span> by{" "}
              <span className="italic">{book.author}</span> ({book.year})
            </div>
            <div className="space-x-2">
              <Link href={`/books/${book.id}`}>
                <button className="bg-green-500 hover:bg-green-700 text-white font-bold py-1 px-3 rounded">
                  View Details{" "}
                </button>
              </Link>

              <button
                onClick={() => handleEdit(book)}
                className="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-1 px-3 rounded"
              >
                Edit
              </button>
              <button
                onClick={() => deleteBook(book.id)}
                className="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-3 rounded"
              >
                Delete
              </button>
            </div>
          </li>
        ))}
      </ul>
      {isModalOpen && (
        <Modal onClose={() => setIsModalOpen(false)}>
          <BookForm book={currentBook} onClose={() => setIsModalOpen(false)} />
        </Modal>
      )}
    </div>
  );
};

export default Dashboard;
