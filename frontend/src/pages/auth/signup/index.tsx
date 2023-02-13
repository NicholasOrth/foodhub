import Router from "next/router"

export default function Register() {

    const handleSubmit = async (e: any) => {
        e.preventDefault();

        const data = {
            name: e.target.name.value,
            email: e.target.email.value,
            password: e.target.password.value,
        };

        if (data.password !== e.target.confirm.value) {
            alert("Passwords do not match");
            return
        }

        const res: Response = await fetch("http://localhost:7100/auth/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        if (res.status === 200) {
            await Router.push("/auth/login");
        }
    }

    return(
        <div className="authform">
            <form onSubmit={handleSubmit}>
                <label htmlFor="name">Full Name</label>
                <input type="text" id="name" name="name" placeholder="Full Name"/>
                <br />

                <label htmlFor="email">email</label>
                <input type="text" id="email" name="email" placeholder="email@whatevermail.com" />
                <br />

                <label htmlFor="password">password</label>
                <input type="password" id="password" name="password" placeholder="**********" />
                <br />

                <label htmlFor="confirm">confirm password</label>
                <input type="password" id="confirm" name="confirm" placeholder="**********" />
                <br />

                <button type="submit">Log In</button>
            </form>
            <br />
            <button onClick={() => Router.push("/auth/login")}>
                Already have an account? Login here.
            </button>
        </div>
    )
}