import React from "react"
import styles from "./home.module.scss"
import { Link } from "react-router-dom"

function Home() {
    return (
        <div className={styles.wrapper}>
            I'm Home Pages.
            <div>使用Link标签</div>
            <Link to="/">回到IndexPage</Link>
        </div>
    )
}

export default Home