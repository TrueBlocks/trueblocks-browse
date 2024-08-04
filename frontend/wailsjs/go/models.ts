export namespace base {
	
	export class Address {
	    address: number[];
	
	    static createFrom(source: any = {}) {
	        return new Address(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = source["address"];
	    }
	}

}

export namespace config {
	
	export class Session {
	    x: number;
	    y: number;
	    width: number;
	    height: number;
	    title: string;
	    lastRoute: string;
	    lastTab: string;
	    lastAddress: string;
	
	    static createFrom(source: any = {}) {
	        return new Session(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.x = source["x"];
	        this.y = source["y"];
	        this.width = source["width"];
	        this.height = source["height"];
	        this.title = source["title"];
	        this.lastRoute = source["lastRoute"];
	        this.lastTab = source["lastTab"];
	        this.lastAddress = source["lastAddress"];
	    }
	}

}

export namespace messages {
	
	export enum Message {
	    COMPLETED = 0,
	    ERROR = 1,
	    WARN = 2,
	    PROGRESS = 3,
	    SERVER = 4,
	    DOCUMENT = 5,
	}
	export class DocumentMsg {
	    filename: string;
	    msg: string;
	
	    static createFrom(source: any = {}) {
	        return new DocumentMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.filename = source["filename"];
	        this.msg = source["msg"];
	    }
	}
	export class ErrorMsg {
	    address: base.Address;
	    errStr: string;
	
	    static createFrom(source: any = {}) {
	        return new ErrorMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.errStr = source["errStr"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ProgressMsg {
	    address: base.Address;
	    have: number;
	    want: number;
	
	    static createFrom(source: any = {}) {
	        return new ProgressMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.have = source["have"];
	        this.want = source["want"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ServerMsg {
	    name: string;
	    message: string;
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new ServerMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.message = source["message"];
	        this.color = source["color"];
	    }
	}

}

export namespace names {
	
	export enum Parts {
	    REGULAR = 2,
	    CUSTOM = 4,
	    PREFUND = 8,
	    BADDRESS = 16,
	}

}

export namespace output {
	
	export class RenderCtx {
	
	
	    static createFrom(source: any = {}) {
	        return new RenderCtx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

export namespace servers {
	
	export enum Type {
	    FILESERVER = 0,
	    SCRAPER = 1,
	    MONITOR = 2,
	    API = 3,
	    IPFS = 4,
	}
	export enum State {
	    STOPPED = 0,
	    RUNNING = 1,
	    PAUSED = 2,
	}
	export class Server {
	    name: string;
	    sleep: number;
	    color: string;
	    // Go type: time
	    started: any;
	    runs: number;
	    state: State;
	
	    static createFrom(source: any = {}) {
	        return new Server(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.sleep = source["sleep"];
	        this.color = source["color"];
	        this.started = this.convertValues(source["started"], null);
	        this.runs = source["runs"];
	        this.state = source["state"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace types {
	
	export class AbiFile {
	    fileSize: number;
	    lastModDate: string;
	    name: string;
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new AbiFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileSize = source["fileSize"];
	        this.lastModDate = source["lastModDate"];
	        this.name = source["name"];
	        this.path = source["path"];
	    }
	}
	export class ChunkStats {
	    addrsPerBlock: number;
	    appsPerAddr: number;
	    appsPerBlock: number;
	    bloomSz: number;
	    chunkSz: number;
	    nAddrs: number;
	    nApps: number;
	    nBlocks: number;
	    nBlooms: number;
	    range: string;
	    rangeEnd: string;
	    ratio: number;
	    recWid: number;
	
	    static createFrom(source: any = {}) {
	        return new ChunkStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.addrsPerBlock = source["addrsPerBlock"];
	        this.appsPerAddr = source["appsPerAddr"];
	        this.appsPerBlock = source["appsPerBlock"];
	        this.bloomSz = source["bloomSz"];
	        this.chunkSz = source["chunkSz"];
	        this.nAddrs = source["nAddrs"];
	        this.nApps = source["nApps"];
	        this.nBlocks = source["nBlocks"];
	        this.nBlooms = source["nBlooms"];
	        this.range = source["range"];
	        this.rangeEnd = source["rangeEnd"];
	        this.ratio = source["ratio"];
	        this.recWid = source["recWid"];
	    }
	}
	export class IndexSummary {
	    addrsPerBlock: number;
	    appsPerAddr: number;
	    appsPerBlock: number;
	    bloomSz: number;
	    chunkSz: number;
	    nAddrs: number;
	    nApps: number;
	    nBlocks: number;
	    nBlooms: number;
	    range: string;
	    rangeEnd: string;
	    ratio: number;
	    recWid: number;
	    chunks: ChunkStats[];
	
	    static createFrom(source: any = {}) {
	        return new IndexSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.addrsPerBlock = source["addrsPerBlock"];
	        this.appsPerAddr = source["appsPerAddr"];
	        this.appsPerBlock = source["appsPerBlock"];
	        this.bloomSz = source["bloomSz"];
	        this.chunkSz = source["chunkSz"];
	        this.nAddrs = source["nAddrs"];
	        this.nApps = source["nApps"];
	        this.nBlocks = source["nBlocks"];
	        this.nBlooms = source["nBlooms"];
	        this.range = source["range"];
	        this.rangeEnd = source["rangeEnd"];
	        this.ratio = source["ratio"];
	        this.recWid = source["recWid"];
	        this.chunks = this.convertValues(source["chunks"], ChunkStats);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ChunkRecord {
	    bloomHash: string;
	    bloomSize: number;
	    indexHash: string;
	    indexSize: number;
	    range: string;
	
	    static createFrom(source: any = {}) {
	        return new ChunkRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bloomHash = source["bloomHash"];
	        this.bloomSize = source["bloomSize"];
	        this.indexHash = source["indexHash"];
	        this.indexSize = source["indexSize"];
	        this.range = source["range"];
	    }
	}
	export class ManifestEx {
	    chain: string;
	    chunks: ChunkRecord[];
	    specification: string;
	    version: string;
	    latestUpdate: string;
	    nBlooms: number;
	    bloomsSize: number;
	    nIndexes: number;
	    indexSize: number;
	
	    static createFrom(source: any = {}) {
	        return new ManifestEx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.chunks = this.convertValues(source["chunks"], ChunkRecord);
	        this.specification = source["specification"];
	        this.version = source["version"];
	        this.latestUpdate = source["latestUpdate"];
	        this.nBlooms = source["nBlooms"];
	        this.bloomsSize = source["bloomsSize"];
	        this.nIndexes = source["nIndexes"];
	        this.indexSize = source["indexSize"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class MonitorEx {
	    address: base.Address;
	    deleted: boolean;
	    ensName: string;
	    fileSize: number;
	    label: string;
	    lastScanned: number;
	    nRecords: number;
	    name: string;
	    transactions: string[];
	
	    static createFrom(source: any = {}) {
	        return new MonitorEx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.deleted = source["deleted"];
	        this.ensName = source["ensName"];
	        this.fileSize = source["fileSize"];
	        this.label = source["label"];
	        this.lastScanned = source["lastScanned"];
	        this.nRecords = source["nRecords"];
	        this.name = source["name"];
	        this.transactions = source["transactions"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class NameEx {
	    address: base.Address;
	    decimals: number;
	    deleted?: boolean;
	    isContract?: boolean;
	    isCustom?: boolean;
	    isErc20?: boolean;
	    isErc721?: boolean;
	    isPrefund?: boolean;
	    name: string;
	    source: string;
	    symbol: string;
	    tags: string;
	    // Go type: base
	    prefund?: any;
	    type: names.Parts;
	
	    static createFrom(source: any = {}) {
	        return new NameEx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.decimals = source["decimals"];
	        this.deleted = source["deleted"];
	        this.isContract = source["isContract"];
	        this.isCustom = source["isCustom"];
	        this.isErc20 = source["isErc20"];
	        this.isErc721 = source["isErc721"];
	        this.isPrefund = source["isPrefund"];
	        this.name = source["name"];
	        this.source = source["source"];
	        this.symbol = source["symbol"];
	        this.tags = source["tags"];
	        this.prefund = this.convertValues(source["prefund"], null);
	        this.type = source["type"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MetaData {
	    client: number;
	    finalized: number;
	    staging: number;
	    ripe: number;
	    unripe: number;
	    chainId?: number;
	    networkId?: number;
	    chain?: string;
	
	    static createFrom(source: any = {}) {
	        return new MetaData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client = source["client"];
	        this.finalized = source["finalized"];
	        this.staging = source["staging"];
	        this.ripe = source["ripe"];
	        this.unripe = source["unripe"];
	        this.chainId = source["chainId"];
	        this.networkId = source["networkId"];
	        this.chain = source["chain"];
	    }
	}
	export class Chain {
	    chain: string;
	    chainId: number;
	    ipfsGateway: string;
	    localExplorer: string;
	    remoteExplorer: string;
	    rpcProvider: string;
	    symbol: string;
	
	    static createFrom(source: any = {}) {
	        return new Chain(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.chainId = source["chainId"];
	        this.ipfsGateway = source["ipfsGateway"];
	        this.localExplorer = source["localExplorer"];
	        this.remoteExplorer = source["remoteExplorer"];
	        this.rpcProvider = source["rpcProvider"];
	        this.symbol = source["symbol"];
	    }
	}
	export class CacheItem {
	    items: any[];
	    lastCached?: string;
	    nFiles: number;
	    nFolders: number;
	    path: string;
	    sizeInBytes: number;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new CacheItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = source["items"];
	        this.lastCached = source["lastCached"];
	        this.nFiles = source["nFiles"];
	        this.nFolders = source["nFolders"];
	        this.path = source["path"];
	        this.sizeInBytes = source["sizeInBytes"];
	        this.type = source["type"];
	    }
	}
	export class StatusEx {
	    cachePath?: string;
	    caches: CacheItem[];
	    chain?: string;
	    chainConfig?: string;
	    chainId?: string;
	    chains: Chain[];
	    clientVersion?: string;
	    hasEsKey?: boolean;
	    hasPinKey?: boolean;
	    indexPath?: string;
	    isApi?: boolean;
	    isArchive?: boolean;
	    isScraping?: boolean;
	    isTesting?: boolean;
	    isTracing?: boolean;
	    networkId?: string;
	    progress?: string;
	    rootConfig?: string;
	    rpcProvider?: string;
	    version?: string;
	    // Go type: MetaData
	    meta?: any;
	    // Go type: MetaData
	    diffs?: any;
	    latestUpdate: string;
	    nFolders: number;
	    nFiles: number;
	    nBytes: number;
	
	    static createFrom(source: any = {}) {
	        return new StatusEx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cachePath = source["cachePath"];
	        this.caches = this.convertValues(source["caches"], CacheItem);
	        this.chain = source["chain"];
	        this.chainConfig = source["chainConfig"];
	        this.chainId = source["chainId"];
	        this.chains = this.convertValues(source["chains"], Chain);
	        this.clientVersion = source["clientVersion"];
	        this.hasEsKey = source["hasEsKey"];
	        this.hasPinKey = source["hasPinKey"];
	        this.indexPath = source["indexPath"];
	        this.isApi = source["isApi"];
	        this.isArchive = source["isArchive"];
	        this.isScraping = source["isScraping"];
	        this.isTesting = source["isTesting"];
	        this.isTracing = source["isTracing"];
	        this.networkId = source["networkId"];
	        this.progress = source["progress"];
	        this.rootConfig = source["rootConfig"];
	        this.rpcProvider = source["rpcProvider"];
	        this.version = source["version"];
	        this.meta = this.convertValues(source["meta"], null);
	        this.diffs = this.convertValues(source["diffs"], null);
	        this.latestUpdate = source["latestUpdate"];
	        this.nFolders = source["nFolders"];
	        this.nFiles = source["nFiles"];
	        this.nBytes = source["nBytes"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TransactionEx {
	    blockNumber: number;
	    date: string;
	    ether: string;
	    from: base.Address;
	    fromName: string;
	    function: string;
	    hasToken: boolean;
	    isError: boolean;
	    logCount: number;
	    to: base.Address;
	    toName: string;
	    transactionIndex: number;
	    // Go type: base
	    wei: any;
	
	    static createFrom(source: any = {}) {
	        return new TransactionEx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.blockNumber = source["blockNumber"];
	        this.date = source["date"];
	        this.ether = source["ether"];
	        this.from = this.convertValues(source["from"], base.Address);
	        this.fromName = source["fromName"];
	        this.function = source["function"];
	        this.hasToken = source["hasToken"];
	        this.isError = source["isError"];
	        this.logCount = source["logCount"];
	        this.to = this.convertValues(source["to"], base.Address);
	        this.toName = source["toName"];
	        this.transactionIndex = source["transactionIndex"];
	        this.wei = this.convertValues(source["wei"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

