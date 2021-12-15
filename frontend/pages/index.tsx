import Head from 'next/head'
import styles from '../styles/Home.module.css'
import {useState} from "react";

export default function Home() {
    const [ url, setUrl ] = useState("");
    const [ isLoading, setLoading ] = useState(false);

    const [ result, setResult ] = useState<null|{url?: string, shortUrl?: string, error?: string}>(null);

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (isLoading) return;

        const putResponse = await fetch(`/api/url`, {
            method: "PUT",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url }),
        }).then(r => r.json());

        if (putResponse.error) {
            setResult({error: putResponse.error});
        } else {
            setResult({url, shortUrl: `${window.location.origin}/${putResponse.hash}`});
        }
    }

    return (
        <div className={styles.container}>
            <Head>
                <title>URL Shorter</title>
            </Head>

            <h1>URL Storter</h1>
            <div>
                <form onSubmit={handleSubmit} className={styles.form}>
                    <input
                        type="text"
                        placeholder="Enter some url"
                        className={styles.input}
                        disabled={isLoading}
                        value={url}
                        onChange={e => setUrl(e.target.value)}
                    />

                    <input
                        type="submit"
                        disabled={isLoading}
                        className={styles.btn}
                        value="Shorten"
                    />
                </form>
            </div>
            <div className={styles.result}>
                {result !== null && result.error && <>
                    error: {result.error}
                </>}

                {result !== null && !result.error && <>
                    {result.url} =&gt; <a href={result.shortUrl}>{result.shortUrl}</a>
                </>}
            </div>
        </div>
    )
}
