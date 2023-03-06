
import Cookies from "universal-cookie";
import {GetServerSidePropsContext} from "next";
import Router from "next/router";
import Navbar from "../../../components/Navbar";

const cookies = new Cookies();

export default function New() {
    const handleSubmit = async (e: any) => {
        e.preventDefault()

        const formData = new FormData()
        formData.append('file', e.target.file.files[0])
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
                await Router.push('/feed')
            }
        } catch (err) {
            console.log(err)
        }
    }

    return (
        <>
            <Navbar />
            <form onSubmit={handleSubmit} encType="multipart/form-data">
                <label htmlFor="caption">Caption</label>
                <input type="text" id="caption" />
                <br />

                <label htmlFor="file">File</label>
                <input type="file" id="file" />
                <br />

                <button type="submit">Submit</button>
            </form>
        </>
    )
}

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
    const jwt = ctx.req.cookies.jwt;

    if (!jwt) {
        return {
            redirect: {
                destination: "/auth/login",
                permanent: false,
            }
        }
    }

    return {
        props: {}
    }
}