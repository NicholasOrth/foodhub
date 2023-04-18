import { useEffect, useState } from "react";
import Navbar from "../../../components/Navbar";

export default function FriendsPage() {
  const [followers, setFollowers] = useState([]);

  useEffect(() => {
    // Make a GET request to the server to retrieve the user's followers
    fetch("http://localhost:7100/user/followers", {
      method: "GET",
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data) => {
        // Update the state with the retrieved followers
        setFollowers(data.followers);
      })
      .catch((error) => {
        console.error("Error fetching followers:", error);
      });
  }, []);

  return (
    <>
      <Navbar />
      <h1>Friends</h1>
      <ul>
        {/* Render the list of followers using JSX */}
        {followers.map((follower) => (
          <li key={follower.id}>{follower.name}</li>
        ))}
      </ul>
    </>
  );
}
