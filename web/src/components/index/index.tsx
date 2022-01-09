import styles from "./index.module.scss"
import { useState } from "react"
function Index() {
    const [count, setCount] = useState(0)

    return (
        <div className={styles.wrapper}>
            Count:{count}
        </div>
    )
}

export default Index;