
import styles from "../src/styles/Navbar.module.css"

const links: {name: string, href: string}[] = [
    {name: "PROFILE", href: "/profile"},
    {name: "FEED", href: "/feed"},
    {name: "NEW", href: "/new"}, 
    {name: "LOGOUT", href: "/auth/logout"},
]

export default function Navbar() {
    return (
        <div className={styles.container}>
            <nav className={styles.nav}>
                {
                    links.map(link => (
                        <a href={link.href} key={link.name}>{link.name}</a>
                    ))
                }
            </nav>
            <div className={styles.spacer} />
        </div>
    )
}