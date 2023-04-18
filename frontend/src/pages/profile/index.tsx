import toast from "react-hot-toast";
import Navbar from "../../../components/Navbar";
import FeedDisplay from "../../../components/FeedDisplay";
import Router from "next/router";
import ProfileView from "../../../components/ProfileView";
import {useEffect, useState} from "react";
import {Me} from "../../../types/Me";

export default function Profile() {
    const [user, setUser] = useState<Me | null>(null);

    useEffect(() => {
        fetch("http://localhost:7100/user/me", {
            method: "GET",
            credentials: "include",
        })
            .then(res => res.json())
            .then(data => {
                console.log(data);
                setUser(data);
            })
    }, []);

    const logout = async () => {
        await fetch("http://localhost:7100/auth/logout", {
            method: "POST",
            credentials: "include",
        })

        toast.success("Logged out successfully");

        await Router.push("/auth/login");
    };

    return (
        <>
            <Navbar />
            {user &&
                <>
                    <ProfileView
                        name={user?.name!}
                        following={user?.following!}
                        followers={user?.followers!}
                        postCount={user?.posts.length!}
                    />
                    <FeedDisplay posts={user?.posts!} />
                </>
            }
            <button onClick={() => { logout(); }}>Logout</button>
        </>
    )
}