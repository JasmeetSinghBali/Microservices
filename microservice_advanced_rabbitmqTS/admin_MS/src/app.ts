import "dotenv/config"
import * as express from 'express'
import * as cors from 'cors'

const app = express()
const PORT = process.env.PORT as string

app.use(express.json())

app.use(cors({
    /**3000 port for REACT version of app
     * 8080 port for VUE version of app
     */
    origin: ['http://localhost:3000','http://localhost:8080']
}))

// home route(test)
app.get('/api/v1',(req:express.Request,res:express.Response)=>{
    try{
      return res.status(200).json({
        status: "Up",
        message: "🦨 API works!!"
      });
    }catch(err:any){
      return res.status(404).json({
        status: "Down",
        message: " 💀 API Broke!!"
      });
    }
});

// minimalistic error handler for no route found
app.use((req:express.Request,res:express.Response)=>{
    return res.status(404).json({
      status: "Failed‍👤",
      message: "Route not found!"
    });
})

app.listen(PORT,()=>console.log(`🐱‍🚀admin_MS running on PORT: ${PORT}`))