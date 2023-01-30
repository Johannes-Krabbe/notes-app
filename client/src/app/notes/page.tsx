import React from "react";
import { request } from "../../helper/apiContext";
import CreateNote from "./CreateNote";
import styles from "./Notes.module.css";

async function getNotes() {
    const res = await request.get("/notes/");
    if (res.status === 200) {
        return res.data;
    } else {
        console.error;
    }
    return undefined;
}

export default async function Home() {
    const notes: any = await getNotes();

    return (
        <main>
            <div>
                {notes?.map((note: any) => {
                    return (
                        <div className={styles.note} key={note.ID.toString()}>
                            <h1>{note.title}</h1>
                            <p>{note.content}</p>
                        </div>
                    );
                })}
            </div>
            <CreateNote />
        </main>
    );
}
