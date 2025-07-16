import React from "react";
import "./Search.css";

interface Props {
  onSearch: (query: string) => void;
  snippets: string[];
  value: string;
}

const Search: React.FC<Props> = ({ onSearch, snippets, value }) => {
  return (
    <div>
      <form>
        <input
          className="search-input"
          value={value}
          onChange={(e) => {
            onSearch(e.target.value);
          }}
          placeholder="Search..."
        />
      </form>
      <div>
        {snippets?.length === 0 ? (
          <p>No results</p>
        ) : (
          snippets?.filter((snip)=>{
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
