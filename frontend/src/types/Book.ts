export interface Book {
    id: number;
    title: string;
    author: string;
    year: number;
    createdAt: string;
    updatedAt: string;
    deletedAt?: string;
    genre?: string;
    isbn?: string;
    publisher?: string;
    description?: string;
}
