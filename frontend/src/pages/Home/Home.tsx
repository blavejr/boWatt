import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getQueryHistory, healthCheck, IQueryHistory, listFiles, queryFile } from "../../api";
import FileList from "../../components/FileList/FileList";
import FileUploader from "../../components/FileUploader/FileUploader";
import Search from "../../components/Search/Search";
import { useDebounce } from "../../utils/useDebounce";
import "./Home.css";

function Home() {
  const [FileHash, setFileHash] = useState<string | null>(null);
  const [healthCheckStatus, setHealthCheckStatus] = useState<string | null>(
    null
  );
  const [snippets, setSnippets] = useState<string[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [userFiles, setUserFiles] = useState<any>([]);
  const [queryHistory, setQueryHistory] = useState<IQueryHistory[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
  const debouncedQuery = useDebounce(searchQuery, 500);

  const navigate = useNavigate()

useEffect(() => {
  const loadData = async () => {
    try {
      const [_, files, history] = await Promise.all([
        healthCheck(),
        listFiles(),
        getQueryHistory(),
      ]);
      setHealthCheckStatus("Backend is healthy");
      setUserFiles(files);
      setQueryHistory(history);
    } catch (err) {
      setHealthCheckStatus("Backend is not healthy");
      setError("Failed to load data. Make sure the backend is running.");
    }
  };
  loadData();
}, []);

useEffect(() => {
  handleSearch(debouncedQuery);
}, [debouncedQuery]);


  const handleUpload = (name: string) => {
    listFiles().then((files) => {
      setUserFiles(files);
    });
  };

  const handleFileSelection = (_fileHash: string) => {
    setError(null)
    setSnippets([])
    setFileHash(_fileHash)
  }

  const handleSearch = async (query: string) => {
    if (!FileHash) {
      setError('Please upload or select a file to search in');
      return;
    }

    if (!query || query.length === 0) {
      setSnippets([]);
      setError(null);
      return;
    }

    try {
      const results = await queryFile(FileHash, query);
      setSnippets(results);
      if (results && results.length === 0) setError("No results found.");
      else setError(null);
    } catch (err) {
      console.error("Search error:", err);
      setError("An error occurred while searching.");
    }
  };


  return (
    <div className="App">
      <header>
        <div className="header-text">
          <h1>Welcome to the App</h1>
          <p>
            This is a simple application that lets you upload a text file and
            search it.
          </p>
          {healthCheckStatus && (
            <p
              className="health-status"
              style={{
                color: healthCheckStatus.includes("not") ? "red" : "green",
              }}
            >
              {healthCheckStatus}
            </p>
          )}
          <p>Selected File: {FileHash?.substring(0, 15) + "..." || "No file selected"}</p>
        </div>
        <FileUploader onUpload={handleUpload} setError={setError} />
        <button className="logout-button" onClick={() => {
          localStorage.removeItem('token');
          navigate('/login')
        }}> Logout </button>
      </header>

      <div className="error-message">{error && <p>{error}</p>}</div>

      <section className="main">
        <div className="search-container">
          <Search onSearch={setSearchQuery} value={searchQuery} snippets={snippets} />
        </div>
        <FileList
          files={userFiles ?? []}
          selectedFile={FileHash}
          onSelect={handleFileSelection}
        />
        <section className="query-history">
          <h1>Query History</h1>
          {queryHistory && queryHistory.length > 0 ? (
            queryHistory
              .filter(hist => !FileHash || hist.FileHash === FileHash)
              .sort((a,b)=>b.Timestamp - a.Timestamp)
              .map((hist, i) => (
                <div className="query-history-entry" key={i} onClick={(e)=>{
                  e.preventDefault()
                  setFileHash(hist?.FileHash)
                  setSearchQuery(hist?.Query)
                }}>
                  <p><strong>File:</strong> {hist?.FileHash?.substring(0, 15)}...</p>
                  <p><strong>Query:</strong> {hist?.Query}</p>
                  <p><strong>When:</strong> {new Date(hist?.Timestamp).toLocaleString()}</p>
                </div>
              ))
          ) : (
            <p>No search history yet.</p>
          )}


        </section>
      </section>
    </div>
  );
}

export default Home;
