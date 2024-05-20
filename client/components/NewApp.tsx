"use client"

import axios from "axios";
import { Loader } from "lucide-react";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { FileUploader } from "react-drag-drop-files";

function NewApp() {
  const [file,setFile] = useState<Blob | null>();
  const [app,setApp] = useState<string>();
  const [branch,setBranch] = useState<string>();
  const [domain,setDomain] = useState<string>();
  const [isLoading,setLoading] = useState(false);
  const formdata = new FormData(); 

  const router = useRouter();

  const handleChange = (file:Blob) => {
    setFile(file);
  }

  const onSubmit = async (e:any) => {
    try{
      e.preventDefault();
      setLoading(true);
      
      const {data} = await axios.post(`http://localhost:8080/api/v1/manual/createApp`,{
        appName:app,
        branchName:branch,
        displayName:domain
      }) 
      console.log(data);

      if(data.status !== "success"){
        console.error("Error Creating App");
        return Promise.reject(data);
      }
    
      formdata.append("file",file!);
      const res = await axios.post(`http://localhost:8080/api/v1/manual/deployApp/${data.message.id}`,formdata,
        {
          headers: {
          'Content-type': "multipart/form-data"
          }
        }
      );
    
      if(res.data.status !== 'success'){
        console.error("Error to start deployment");
        return Promise.reject(data);
      }
      setApp("");
      setBranch("");
      setDomain("");
      setFile(null);

      setLoading(false);
      router.refresh();

      const getStatus = await axios.get(`http://localhost:8080/api/v1/manual/getStatus/${data.message.id}`);

      if(getStatus.data.status !== 'success'){
        console.error("Error to get job status");
        return Promise.reject(data);
      }
      router.refresh();
    }catch(err){
      console.error(err);
      return Promise.reject(err);
    }finally{
      setLoading(false);
    }
  }

  return (
    <div className="">
      <form  className="grid grid-cols-2 justify-center gap-8 " onSubmit={onSubmit}>
        <div className="flex flex-col gap-1">
          <label htmlFor="app">Enter App Name</label>
          <input 
            name="app" 
            type="text" 
            placeholder="Landing Page, Portfolio" 
            value={app}
            required 
            onChange={(e) => setApp(e.target.value)}
            className="p-3 rounded-md bg-slate-700 placeholder:text-sm placeholder-slate-400 outline-none focus:text-slate-300"
          />
        </div>

        <div className="flex flex-col gap-1">
          <label htmlFor="branch">Enter Branch Name</label>
          <input 
            name="branch" 
            type="text" 
            placeholder="main, dev, prod, test" 
            value={branch}
            required
            onChange={(e) => setBranch(e.target.value)}
            className="p-3 rounded-md bg-slate-700 placeholder:text-sm placeholder-slate-400 outline-none focus:text-slate-300"
          />
        </div>
        
        <div className="flex flex-col gap-1">
          <label htmlFor="domain">Enter Domain Name</label>
          <input 
            name="domain" 
            type="text" 
            placeholder="google" 
            value={domain}
            required
            onChange={(e) => setDomain(e.target.value)}
            className="p-3 rounded-md bg-slate-700 placeholder:text-sm placeholder-slate-400 outline-none focus:text-slate-300"
          />
        </div>

        <div className="flex flex-col gap-1 overflow-hidden whitespace-pre">
          <label htmlFor="zipFile">Upload Zip file</label>
          <FileUploader 
            handleChange={handleChange} 
            name="zipFile" 
            types={["zip"]} 
            label={"Upload or drop a zip file right here"}
            required
            fileOrFiles={file}
            />
        </div>
        
        <div className="flex gap-6 items-center">
          <button
             type="submit" 
             value="Submit" 
             className="border-2 border-slate-200 rounded-md py-2 px-4 hover:bg-slate-800 hover:text-slate-100 cursor-pointer"
          >{
            !isLoading ? "Submit" :(
              <Loader className="h-6 w-6 animate-spin"/>
            )
          }</button>
        </div>
      </form>
    </div>
  );
}

export default NewApp;