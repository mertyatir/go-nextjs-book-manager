import { Book } from "@/types/Book";

const BASE_URL = "http://localhost:8000";

export async function getBooks() {
    const response = await fetch(`${BASE_URL}/books`);
    if (!response.ok) {
        throw new Error("Failed to fetch books");
    }
    return response.json();

}

export async function postBook(book: Book) {
    const response = await fetch(`${BASE_URL}/books`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(book),
    });

    if (!response.ok) {
        throw new Error("Failed to add book");
    }
    return response.json();
}

export async function editBook(updatedBook: Book) {
    const response = await fetch(
        `${BASE_URL}/books/${updatedBook.id}`,
        {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(updatedBook),
        }
    );

    if (!response.ok) {
        throw new Error("Failed to update book");
    }

    return response.json();
}



export async function deleteBook(id: number) {
    const response = await fetch(`${BASE_URL}/books/${id}`, {
        method: "DELETE",
    });

    if (!response.ok) {
        throw new Error("Failed to delete book");
    }

}
