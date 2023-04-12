
import Cookies from "universal-cookie";
import {GetServerSidePropsContext} from "next";
import Router from "next/router";
import Navbar from "../../../components/Navbar";
import styles from "@/styles/New.module.css";
import toast from "react-hot-toast";
const cookies = new Cookies();

export default function New() {
    const acceptableTypes = ["image/png", "image/jpeg", "image/jpg"];
    const handleSubmit = async (e: any) => {
        e.preventDefault()
        
        const file = e.target.file.files[0];
        if(acceptableTypes.includes(file.type) === false) {
            toast.error("File type not supported");
            return;
        }

        const formData = new FormData()
        formData.append('file', file)
        formData.append('caption', e.target.caption.value)

        const jwt = cookies.get('jwt');

        try {
            const res = await fetch('http://localhost:7100/post/create', {
                method: 'POST',
                headers: {
                    'Cookie': "jwt=" + jwt
                },
                body: formData,
                credentials: 'include'
            })

            const data = await res.json()

            if (res.ok) {
                toast.success('Post posted!')
                await Router.push('/feed')
            }
            
        } catch (err) {
            toast.error("Issue posting post")
        }
    }

    return (
        <>
            <Navbar />
            <div className={styles.container}>
                <h1>New Post</h1>
                <form onSubmit={handleSubmit}>
                    <label>
                        caption
                        <textarea id="caption" name="caption" placeholder="caption"></textarea>
                    </label>

                    <label>
                        Photo
                        <input type="file" id="file" name="file" />
                    </label>

                    <button type="submit">Post</button>
                </form>
            </div>
        </>
    )
}

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
    const jwt = ctx.req.cookies.jwt;

    if (!jwt) {
        return {
            redirect: {
                destination: "/",
                permanent: false,
            }
        }
    }

    return {
        props: {}
    }
}