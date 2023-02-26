
import Cookies from "universal-cookie";

const cookies = new Cookies();

export default function New() {
    const handleSubmit = async (e: any) => {
        e.preventDefault()

        const formData = new FormData()
        formData.append('file', e.target.file.files[0])
        formData.append('caption', 'test')

        const jwt = cookies.get('jwt');

        try {
            const res = await fetch('http://localhost:7100/image/upload', {
                method: 'POST',
                headers: {
                    'Cookie': "jwt=" + jwt
                },
                body: formData,
                credentials: 'include'
            })

            const data = res.json()
            data.then((data: any) => {
                console.log(data);
            }).catch((err: any) => {
                console.log("");
            })
        } catch (err) {
            console.log(err)
        }
    }

    return (
        <form onSubmit={handleSubmit} encType="multipart/form-data">
            <input type="file" id="file" />
            <button type="submit">Submit</button>
        </form>
    )
}
