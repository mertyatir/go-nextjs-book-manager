import React, { useContext, useState, useEffect } from "react";
import { BookContext } from "../context/BookContext";
import { Book } from "@/types/Book";

interface BookFormProps {
  book?: Book;
  onClose: () => void;
}

const BookForm: React.FC<BookFormProps> = ({ book, onClose }) => {
  const [title, setTitle] = useState("");
  const [author, setAuthor] = useState("");
  const [year, setYear] = useState("");
  const [genre, setGenre] = useState("");
  const [isbn, setIsbn] = useState("");
  const [publisher, setPublisher] = useState("");
  const [description, setDescription] = useState("");

  useEffect(() => {
    if (book) {
      setTitle(book.title);
      setAuthor(book.author);
      setYear(book.year.toString());

      setGenre(book.genre ?? "");
      setIsbn(book.isbn ?? "");
      setPublisher(book.publisher ?? "");
      setDescription(book.description ?? "");
    }
  }, [book]);

  const context = useContext(BookContext);

  if (!context) {
    return <div>Context not available</div>;
  }

  const { addBook, editBook } = context;

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const newBook: Book = {
      id: book ? book.id : Date.now(),
      title,
      author,
      year: parseInt(year),
      createdAt: book ? book.createdAt : new Date().toISOString(),
      updatedAt: new Date().toISOString(),
      genre,
      isbn,
      publisher,
      description,
    };
    if (book) {
      editBook(newBook);
    } else {
      addBook(newBook);
    }
    onClose();
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="max-w-lg mx-auto p-6 bg-white shadow-md rounded-lg"
    >
      <h2 className="text-2xl font-bold mb-4">Add a New Book</h2>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="title"
        >
          Title
        </label>
        <input
          type="text"
          id="title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="author"
        >
          Author
        </label>
        <input
          type="text"
          id="author"
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
          required
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="year"
        >
          Year
        </label>
        <input
          type="number"
          id="year"
          value={year}
          onChange={(e) => setYear(e.target.value)}
          required
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="genre"
        >
          Genre
        </label>
        <input
          type="text"
          id="genre"
          value={genre}
          onChange={(e) => setGenre(e.target.value)}
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="isbn"
        >
          ISBN
        </label>
        <input
          type="text"
          id="isbn"
          value={isbn}
          onChange={(e) => setIsbn(e.target.value)}
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="publisher"
        >
          Publisher
        </label>
        <input
          type="text"
          id="publisher"
          value={publisher}
          onChange={(e) => setPublisher(e.target.value)}
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="description"
        >
          Description
        </label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        />
      </div>
      <button
        type="submit"
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      >
        Submit
      </button>
    </form>
  );
};

export default BookForm;
