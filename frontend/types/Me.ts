import {Post} from "./Post";

export type Me = {
    name: string;
    posts: Post[];
    followers: number;
    following: number;
}