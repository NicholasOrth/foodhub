import {GetServerSidePropsContext} from "next";

export default function Profile(
    props: {data: {email: string, name: string}})
{
    return (
        <div>
            <h1>Profile</h1>
            <p>{props.data.name}</p>
            <p>{props.data.email}</p>
        </div>
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

    const res: Response = await fetch('http://localhost:7100/user', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Cookie': "jwt=" + jwt,
        },
        credentials: 'include',
    })

    const data: {email: string, name: string} = await res.json();

    return {
        props: {
            data: {
                email: data.email,
                name: data.name,
            }
        }
    }
}
