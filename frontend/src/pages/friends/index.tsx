import Navbar from "../../../components/Navbar";
import {Profile} from "../../../types/Profile";
import {useState} from "react";
import ProfileView from "../../../components/ProfileView";

import styles from "../../styles/Friends.module.css";

export default function FriendsPage(props: any) {

    const [profileList, setProfileList] = useState<Profile[]>([]);

    const handleSubmit = async (e: any) => {
        e.preventDefault();

        setProfileList([])

        const query = e.target.query.value;

        const res = await fetch(`http://localhost:7100/user/search/${query}`);

        if (res.ok) {
            const data: { users: number[] } = await res.json();

            for (const id of data.users) {
                const res = await fetch(`http://localhost:7100/user/${id}`);
                const data: Profile = await res.json();

                console.log(data);

                setProfileList(profileList => [...profileList, data]);
            }
        }
    }

    return (
        <>
            <Navbar />
            <h1>Add Friend</h1>

            <form onSubmit={handleSubmit}>
                <label>
                    Search:
                    <input type="text" name="query" />
                </label>
                <input type="submit" value="Submit" />
            </form>

            <ul className={styles.profileList}>
                {
                    profileList.map(profile => {
                       return (
                              <li key={profile.id}>
                                  <ProfileView
                                      name={profile.name}
                                      followers={profile.followers}
                                      following={profile.following}
                                      postCount={profile.posts}
                                  />
                              </li>
                       )
                    })
                }
            </ul>
        </>
    )
}