import { useState, useEffect} from "react";
import Navbar from "../../../components/Navbar";
import styles from "../../../src/styles/Friends.module.css";
import { userInfo } from "os";

export default function FriendsPage(props: any) {
    //goals for this page:
//have a search bar that finds users
//have a list of friends
//be able to add friends
//be able to remove friends

const handleSubmit = async (e: any) => { //search for a friend
    e.preventDefault(); //prevent page from refreshing
   const friendQuerey = e.target.friend.value 

    try {
        const response = await fetch(`http://localhost:7100/user/search/${friendQuerey}`) 

        if(response.ok){
            const data: number[] = await response.json() //parse the data from the response
            data.forEach(async (id: number) => {
                const user = await fetch(`http://localhost:7100/user/${id}`) //get the user info
                const userInfo = await user.json()
            }) 
        }
    }
      catch (e: any) {
        console.log(e.message)
    }
}
    





    return (
        <>
            <Navbar />
            <h1>Friends Search: </h1>
            <form onSubmit={handleSubmit}> 
                <label>
                    Search for a friend
                    <input type="text" id="friend" name="friend" placeholder="friend" />
                </label>
                <button type="submit">Search</button>
            </form>

            <ul>
                

        </>
    )
}