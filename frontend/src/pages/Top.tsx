import React, { useState } from 'react'
import axios from 'axios'



const Top = () => {
    const [status, setStatus] = useState("first step");


    const getXemBalance = () => {
        axios.get("http://localhost:8080/xem")
          .then((res: any) => {
              console.log(res.data)
              setStatus(res.data)
          })
          .catch(
            (err: any) => {console.log("レスポンスに発生しました:", err)}
          )
      }

    return (
        <div>
            <h4>{status}</h4>   
            <button onClick={getXemBalance}>XENの情報を取得する</button>
        </div>
    )
}

export default Top