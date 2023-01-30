"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { request } from "../../helper/apiContext";

export default function CreateNote() {
    const [title, setTitle] = useState<string>("");
    const [content, setContent] = useState<string>("");

    const router = useRouter();

    const create = async (e: any) => {
        e.preventDefault();
        try {
            const res = await request.post("/notes/", {
                title: "test",
                content: "tseartie",
            });
            setContent("");
            setTitle("");
            // router.refresh();
        } catch (e) {
            console.error(e);
        }
    };

    return (
        <>
            <button onClick={create}>test</button>
            <form onSubmit={create}>
                <h3>Create a new Note</h3>
                <input
                    type="text"
                    placeholder="Title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
                <textarea
                    placeholder="Content"
                    value={content}
                    onChange={(e) => setContent(e.target.value)}
                />
                <input type="submit" value="create Note" />
            </form>
        </>
    );
}
