"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { request } from "../../helper/apiContext";

async function createNotes() {
    console.log("hello");
    const res = await request.post("/notes/", {
        title: "Test Note",
        content: "This is some content",
    });
}

export default function CreateNote() {
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");

    const router = useRouter();

    const create = async () => {
        try {
            const res = await request.post("/notes/", {
                title: "test",
                content: "content",
            });
            setContent("");
            setTitle("");
            router.refresh();
        } catch (e) {
            console.error(e);
        }
    };

    return (
        <>
            <button onClick={createNotes}>test</button>
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
                <button type="submit">Create note</button>
            </form>
        </>
    );
}
