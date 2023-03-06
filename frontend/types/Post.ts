export type Post = {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;

    username: string;
    caption: string;
    imgPath: string;
    likes: number;
}