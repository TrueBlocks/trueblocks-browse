#### wizard-index.md

The Wizard Index page checks your local computer's hard drive for a downloaded copy of chunks from the Unchained Index.

This process may take several minutes (or more than an hour depending on your Internet connection).

Downloading the index portions of the Unchained Index can be delayed if you so choose, however, the performance of the system is significantly faster if you download them all prior to using TrueBlocks Browse.

If you choose not to download the index portions, they will be downloaded "as needed." For example, if you were to query for the history of `trueblocks.eth`, then only those portions of the index that hit on the chunk's bloom filter will be downloaded. This greatly lessens the size of the Unchained Index on your hard drive, but it does increase the download speed of an address's history the first time it's accessed. (Second and subsequent queries are as fast as the full download mode.)
