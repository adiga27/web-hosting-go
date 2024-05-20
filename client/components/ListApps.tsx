"use client"
import Link from "next/link"

import { Loader, LucideLink, Wrench } from 'lucide-react';
import axios from "axios";
import { useRouter } from "next/navigation";
import { useState } from "react";

type ListApps = {
    id:string,
    appName:string
    appId?:string
    branchName:string
    url:string
    jobId?:string
    status?:string
    displayName:string
}
type ListAppsProps ={
    apps: ListApps[]
}

function ListApps(data:ListAppsProps) {
    const router = useRouter();
    const [key,setKey] = useState("");
    const onDelete = async (e:any,id:string) => {
        try{
            setKey(id);
            e.preventDefault();
            const {data} = await axios.delete(`http://localhost:8080/api/v1/manual/deleteApp/${id}`);
            if(data.status != 'success'){
                console.log('error');
                return Promise.reject(data);
            }
            setKey("");
            router.refresh();
        }catch(err){
            console.log('error');
            return Promise.reject(err);
        }finally{
            setKey("");
        }
    }
    return (
        <>
            <h2 className="text-2xl">Apps</h2>
            <div className="grid grid-cols-2 gap-8">
                {data.apps?.map(app =>
                    <div key={app.id} className=" border-2 border-zinc-600 rounded-md p-4 overflow-hidden whitespace-pre">
                        <div className="flex justify-between items-center mb-3">
                            <h2 className="text-xl text-zinc-50">{app.appName}</h2>
                            <div 
                                className={`px-2 py-0.5 text-zinc-50 w-min text-xs rounded-md
                                ${app.status === 'SUCCEED' ? 'bg-green-700' : app.status === 'FAILED' ? "bg-red-700" : "bg-yellow-700"  }`}
                            >{app.status ? app.status : "NOT UPLOADED"}</div>
                        </div>
                        <div className="">
                            <Wrench className="h-4 w-4 inline-block mr-2 stroke-zinc-300"/>
                            <p className="inline-block text-zinc-300">{app.branchName}</p>
                        </div>
                        <div>
                            <LucideLink className="h-4 w-4 inline-block mr-2 stroke-slate-300"/>
                            <p className="text-zinc-300 inline-block text-ellipsis">
                                url:
                                <Link 
                                    href={app.url} 
                                    target="_blank"
                                    className="hover:text-zinc-200 "
                                > {app.url}</Link>
                            </p>
                        </div>
                        <div className="flex justify-end mt-4 text-zinc-200">
                            <button 
                                className="cursor-pointer bg-red-700 py-1.5 px-3 hover:bg-red-800 rounded-md"
                                disabled={app.id === key}
                                onClick={(e:any) => onDelete(e,app.id)}
                            >{app.id !== key ? "Delete App" : <Loader className="h-6 w-6 animate-spin"/>}</button>
                        </div>
                    </div>
                )}             
            </div>
        </>
    );
}

export default ListApps;