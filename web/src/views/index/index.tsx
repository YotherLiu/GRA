import styles from "./index.module.scss"
import { Link } from "react-router-dom"

function Index() {
    return (
        <div className={styles.wrapper}>
            <div>☺Hello,React Template</div>
            I'm Index Pages.
            <div>使用Link标签跳转</div>
            <Link to="/home">Go Home Page.</Link>
        </div>
    )
}

export default Index