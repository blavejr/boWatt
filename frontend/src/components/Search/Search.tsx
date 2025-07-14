import React, { useState } from "react";
import "./Search.css";

interface Props {
  onSearch: (query: string) => void;
  snippets: string[];
}

const Search: React.FC<Props> = ({ onSearch, snippets }) => {
  const [query, setQuery] = useState("");

  return (
    <div>
      <form>
        <input
          className="search-input"
          value={query}
          onChange={(e) => {
            setQuery(e.target.value);
            onSearch(e.target.value);
          }}
          placeholder="Search..."
        />
      </form>
      <div>
        {snippets.length === 0 ? (
          <p>No results</p>
        ) : (
          snippets.filter((snip)=>{
            return snip.trim() !== ""
          }).map((snippet, _index) => (
            <div className="snippet" key={_index}>
              {snippet}
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default Search;
