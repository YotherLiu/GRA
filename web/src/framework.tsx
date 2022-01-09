// Framework

let Framework = function (props: any) {
    return (
        <div>
            {/* <div>我是Framework(头部)页面！</div> */}
            {props.children}
            {/* <div>我是Framework(脚部)页面！</div> */}
        </div>
    )
}

export default Framework