import Router from "./routers/route"
import Framework from "./framework"

function App() {
  return (
    <>
      <div id="wrapper">
        <Framework>
          <Router />
        </Framework>
      </div>
    </>
  )
}

export default App