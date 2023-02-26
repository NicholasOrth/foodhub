import Cookies from "universal-cookie"

const cookies = new Cookies();

export function printString(str: string | null) {
    console.log(str)
}

export function setCookie(name: string, value: string) {
    cookies.set(name, value, {
        path: "/",
        maxAge: 3600,
        sameSite: "strict",
        secure: false,
        domain: "localhost",
        httpOnly: false,
    });
}