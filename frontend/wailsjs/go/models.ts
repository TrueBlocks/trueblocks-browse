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
	export class Hash {
	    hash: number[];
	
	    static createFrom(source: any = {}) {
	        return new Hash(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hash = source["hash"];
	    }
	}

}

export namespace config {
	
	export class Daemons {
	    freshen: boolean;
	    scraper: boolean;
	    ipfs: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Daemons(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.freshen = source["freshen"];
	        this.scraper = source["scraper"];
	        this.ipfs = source["ipfs"];
	    }
	}
	export class Session {
	    x: number;
	    y: number;
	    width: number;
	    height: number;
	    title: string;
	    lastRoute: string;
	    lastSub: {[key: string]: string};
	    lastHelp: boolean;
	    daemons: Daemons;
	
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
	        this.lastSub = source["lastSub"];
	        this.lastHelp = source["lastHelp"];
	        this.daemons = this.convertValues(source["daemons"], Daemons);
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

export namespace daemons {
	
	export enum State {
	    STOPPED = "Stopped",
	    RUNNING = "Running",
	    PAUSED = "Paused",
	}
	export class Daemon {
	    name: string;
	    sleep: number;
	    color: string;
	    // Go type: time
	    started: any;
	    ticks: number;
	    state: State;
	
	    static createFrom(source: any = {}) {
	        return new Daemon(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.sleep = source["sleep"];
	        this.color = source["color"];
	        this.started = this.convertValues(source["started"], null);
	        this.ticks = source["ticks"];
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

export namespace messages {
	
	export enum Message {
	    COMPLETED = "Completed",
	    ERROR = "Error",
	    WARN = "Warn",
	    PROGRESS = "Progress",
	    DAEMON = "Daemon",
	    DOCUMENT = "Document",
	    NAVIGATE = "Navigate",
	    TOGGLEHELP = "ToggleHelp",
	}
	export class DaemonMsg {
	    name: string;
	    message: string;
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new DaemonMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.message = source["message"];
	        this.color = source["color"];
	    }
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
	
	export class NavigateMsg {
	    route: string;
	
	    static createFrom(source: any = {}) {
	        return new NavigateMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.route = source["route"];
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

export namespace types {
	
	export class Parameter {
	    components?: Parameter[];
	    indexed?: boolean;
	    internalType?: string;
	    name: string;
	    strDefault?: string;
	    type: string;
	    value?: any;
	
	    static createFrom(source: any = {}) {
	        return new Parameter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.components = this.convertValues(source["components"], Parameter);
	        this.indexed = source["indexed"];
	        this.internalType = source["internalType"];
	        this.name = source["name"];
	        this.strDefault = source["strDefault"];
	        this.type = source["type"];
	        this.value = source["value"];
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
	export class Function {
	    anonymous?: boolean;
	    constant?: boolean;
	    encoding: string;
	    inputs: Parameter[];
	    message?: string;
	    name: string;
	    outputs: Parameter[];
	    signature?: string;
	    stateMutability?: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new Function(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.anonymous = source["anonymous"];
	        this.constant = source["constant"];
	        this.encoding = source["encoding"];
	        this.inputs = this.convertValues(source["inputs"], Parameter);
	        this.message = source["message"];
	        this.name = source["name"];
	        this.outputs = this.convertValues(source["outputs"], Parameter);
	        this.signature = source["signature"];
	        this.stateMutability = source["stateMutability"];
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
	export class Abi {
	    address: base.Address;
	    fileSize: number;
	    functions: Function[];
	    isKnown: boolean;
	    lastModDate: string;
	    nEvents: number;
	    nFunctions: number;
	    name: string;
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new Abi(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.fileSize = source["fileSize"];
	        this.functions = this.convertValues(source["functions"], Function);
	        this.isKnown = source["isKnown"];
	        this.lastModDate = source["lastModDate"];
	        this.nEvents = source["nEvents"];
	        this.nFunctions = source["nFunctions"];
	        this.name = source["name"];
	        this.path = source["path"];
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
	export class AbiContainer {
	    address: base.Address;
	    fileSize: number;
	    functions: Function[];
	    isKnown: boolean;
	    lastModDate: string;
	    nEvents: number;
	    nFunctions: number;
	    name: string;
	    path: string;
	    items: Abi[];
	    nItems: number;
	    largestFile: string;
	    mostFunctions: string;
	    mostEvents: string;
	
	    static createFrom(source: any = {}) {
	        return new AbiContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.fileSize = source["fileSize"];
	        this.functions = this.convertValues(source["functions"], Function);
	        this.isKnown = source["isKnown"];
	        this.lastModDate = source["lastModDate"];
	        this.nEvents = source["nEvents"];
	        this.nFunctions = source["nFunctions"];
	        this.name = source["name"];
	        this.path = source["path"];
	        this.items = this.convertValues(source["items"], Abi);
	        this.nItems = source["nItems"];
	        this.largestFile = source["largestFile"];
	        this.mostFunctions = source["mostFunctions"];
	        this.mostEvents = source["mostEvents"];
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
	
	export class IndexContainer {
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
	    items: ChunkStats[];
	    nItems: number;
	
	    static createFrom(source: any = {}) {
	        return new IndexContainer(source);
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
	        this.items = this.convertValues(source["items"], ChunkStats);
	        this.nItems = source["nItems"];
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
	export class Log {
	    address: base.Address;
	    articulatedLog?: Function;
	    blockHash: base.Hash;
	    blockNumber: number;
	    data?: string;
	    logIndex: number;
	    timestamp?: number;
	    topics?: base.Hash[];
	    transactionHash: base.Hash;
	    transactionIndex: number;
	
	    static createFrom(source: any = {}) {
	        return new Log(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.articulatedLog = this.convertValues(source["articulatedLog"], Function);
	        this.blockHash = this.convertValues(source["blockHash"], base.Hash);
	        this.blockNumber = source["blockNumber"];
	        this.data = source["data"];
	        this.logIndex = source["logIndex"];
	        this.timestamp = source["timestamp"];
	        this.topics = this.convertValues(source["topics"], base.Hash);
	        this.transactionHash = this.convertValues(source["transactionHash"], base.Hash);
	        this.transactionIndex = source["transactionIndex"];
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
	export class ManifestContainer {
	    chain: string;
	    chunks: ChunkRecord[];
	    specification: string;
	    version: string;
	    items: ChunkRecord[];
	    nItems: number;
	    latestUpdate: string;
	    nBlooms: number;
	    bloomsSize: number;
	    nIndexes: number;
	    indexSize: number;
	
	    static createFrom(source: any = {}) {
	        return new ManifestContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.chunks = this.convertValues(source["chunks"], ChunkRecord);
	        this.specification = source["specification"];
	        this.version = source["version"];
	        this.items = this.convertValues(source["items"], ChunkRecord);
	        this.nItems = source["nItems"];
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
	
	export class Monitor {
	    address: base.Address;
	    deleted: boolean;
	    fileSize: number;
	    lastScanned: number;
	    nRecords: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Monitor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.deleted = source["deleted"];
	        this.fileSize = source["fileSize"];
	        this.lastScanned = source["lastScanned"];
	        this.nRecords = source["nRecords"];
	        this.name = source["name"];
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
	export class MonitorContainer {
	    address: base.Address;
	    deleted: boolean;
	    fileSize: number;
	    lastScanned: number;
	    nRecords: number;
	    name: string;
	    items: Monitor[];
	    nItems: number;
	    nNamed: number;
	    nDeleted: number;
	    monitorMap: {[key: string]: Monitor};
	
	    static createFrom(source: any = {}) {
	        return new MonitorContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.deleted = source["deleted"];
	        this.fileSize = source["fileSize"];
	        this.lastScanned = source["lastScanned"];
	        this.nRecords = source["nRecords"];
	        this.name = source["name"];
	        this.items = this.convertValues(source["items"], Monitor);
	        this.nItems = source["nItems"];
	        this.nNamed = source["nNamed"];
	        this.nDeleted = source["nDeleted"];
	        this.monitorMap = this.convertValues(source["monitorMap"], Monitor, true);
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
	export class Name {
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
	    parts?: number;
	
	    static createFrom(source: any = {}) {
	        return new Name(source);
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
	        this.parts = source["parts"];
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
	export class NameContainer {
	    names: Name[];
	    namesMap: {[key: string]: Name};
	    nItems: number;
	    nContracts: number;
	    nErc20s: number;
	    nErc721s: number;
	    nCustom: number;
	    nRegular: number;
	    nPrefund: number;
	    nBaddress: number;
	    nDeleted: number;
	
	    static createFrom(source: any = {}) {
	        return new NameContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.names = this.convertValues(source["names"], Name);
	        this.namesMap = this.convertValues(source["namesMap"], Name, true);
	        this.nItems = source["nItems"];
	        this.nContracts = source["nContracts"];
	        this.nErc20s = source["nErc20s"];
	        this.nErc721s = source["nErc721s"];
	        this.nCustom = source["nCustom"];
	        this.nRegular = source["nRegular"];
	        this.nPrefund = source["nPrefund"];
	        this.nBaddress = source["nBaddress"];
	        this.nDeleted = source["nDeleted"];
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
	
	export class Receipt {
	    blockHash?: base.Hash;
	    blockNumber: number;
	    contractAddress?: base.Address;
	    cumulativeGasUsed?: number;
	    effectiveGasPrice?: number;
	    from?: base.Address;
	    gasUsed: number;
	    isError?: boolean;
	    logs: Log[];
	    status: number;
	    to?: base.Address;
	    transactionHash: base.Hash;
	    transactionIndex: number;
	
	    static createFrom(source: any = {}) {
	        return new Receipt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.blockHash = this.convertValues(source["blockHash"], base.Hash);
	        this.blockNumber = source["blockNumber"];
	        this.contractAddress = this.convertValues(source["contractAddress"], base.Address);
	        this.cumulativeGasUsed = source["cumulativeGasUsed"];
	        this.effectiveGasPrice = source["effectiveGasPrice"];
	        this.from = this.convertValues(source["from"], base.Address);
	        this.gasUsed = source["gasUsed"];
	        this.isError = source["isError"];
	        this.logs = this.convertValues(source["logs"], Log);
	        this.status = source["status"];
	        this.to = this.convertValues(source["to"], base.Address);
	        this.transactionHash = this.convertValues(source["transactionHash"], base.Hash);
	        this.transactionIndex = source["transactionIndex"];
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
	export class Rewards {
	    // Go type: base
	    block: any;
	    // Go type: base
	    nephew: any;
	    // Go type: base
	    txFee: any;
	    // Go type: base
	    uncle: any;
	
	    static createFrom(source: any = {}) {
	        return new Rewards(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.block = this.convertValues(source["block"], null);
	        this.nephew = this.convertValues(source["nephew"], null);
	        this.txFee = this.convertValues(source["txFee"], null);
	        this.uncle = this.convertValues(source["uncle"], null);
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
	export class StatusContainer {
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
	    items: CacheItem[];
	    nItems: number;
	    latestUpdate: string;
	    nFolders: number;
	    nFiles: number;
	    nBytes: number;
	
	    static createFrom(source: any = {}) {
	        return new StatusContainer(source);
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
	        this.items = this.convertValues(source["items"], CacheItem);
	        this.nItems = source["nItems"];
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
	export class TraceResult {
	    address?: base.Address;
	    code?: string;
	    gasUsed?: number;
	    output?: string;
	
	    static createFrom(source: any = {}) {
	        return new TraceResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.code = source["code"];
	        this.gasUsed = source["gasUsed"];
	        this.output = source["output"];
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
	export class TraceAction {
	    address?: base.Address;
	    author?: base.Address;
	    // Go type: base
	    balance?: any;
	    callType: string;
	    from: base.Address;
	    gas: number;
	    init?: string;
	    input?: string;
	    refundAddress?: base.Address;
	    rewardType?: string;
	    selfDestructed?: base.Address;
	    to: base.Address;
	    // Go type: base
	    value: any;
	
	    static createFrom(source: any = {}) {
	        return new TraceAction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.author = this.convertValues(source["author"], base.Address);
	        this.balance = this.convertValues(source["balance"], null);
	        this.callType = source["callType"];
	        this.from = this.convertValues(source["from"], base.Address);
	        this.gas = source["gas"];
	        this.init = source["init"];
	        this.input = source["input"];
	        this.refundAddress = this.convertValues(source["refundAddress"], base.Address);
	        this.rewardType = source["rewardType"];
	        this.selfDestructed = this.convertValues(source["selfDestructed"], base.Address);
	        this.to = this.convertValues(source["to"], base.Address);
	        this.value = this.convertValues(source["value"], null);
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
	export class Trace {
	    action?: TraceAction;
	    articulatedTrace?: Function;
	    blockHash: base.Hash;
	    blockNumber: number;
	    error?: string;
	    result?: TraceResult;
	    subtraces: number;
	    timestamp: number;
	    traceAddress: number[];
	    transactionHash: base.Hash;
	    transactionIndex: number;
	    type?: string;
	    transactionPosition?: number;
	
	    static createFrom(source: any = {}) {
	        return new Trace(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.action = this.convertValues(source["action"], TraceAction);
	        this.articulatedTrace = this.convertValues(source["articulatedTrace"], Function);
	        this.blockHash = this.convertValues(source["blockHash"], base.Hash);
	        this.blockNumber = source["blockNumber"];
	        this.error = source["error"];
	        this.result = this.convertValues(source["result"], TraceResult);
	        this.subtraces = source["subtraces"];
	        this.timestamp = source["timestamp"];
	        this.traceAddress = source["traceAddress"];
	        this.transactionHash = this.convertValues(source["transactionHash"], base.Hash);
	        this.transactionIndex = source["transactionIndex"];
	        this.type = source["type"];
	        this.transactionPosition = source["transactionPosition"];
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
	
	
	export class Statement {
	    accountedFor: base.Address;
	    // Go type: base
	    amountIn?: any;
	    // Go type: base
	    amountOut?: any;
	    assetAddr: base.Address;
	    assetSymbol: string;
	    // Go type: base
	    begBal: any;
	    blockNumber: number;
	    // Go type: base
	    correctingIn?: any;
	    // Go type: base
	    correctingOut?: any;
	    correctingReason?: string;
	    decimals: number;
	    // Go type: base
	    endBal: any;
	    // Go type: base
	    gasOut?: any;
	    // Go type: base
	    internalIn?: any;
	    // Go type: base
	    internalOut?: any;
	    logIndex: number;
	    // Go type: base
	    minerBaseRewardIn?: any;
	    // Go type: base
	    minerNephewRewardIn?: any;
	    // Go type: base
	    minerTxFeeIn?: any;
	    // Go type: base
	    minerUncleRewardIn?: any;
	    // Go type: base
	    prefundIn?: any;
	    // Go type: base
	    prevBal?: any;
	    priceSource: string;
	    recipient: base.Address;
	    // Go type: base
	    selfDestructIn?: any;
	    // Go type: base
	    selfDestructOut?: any;
	    sender: base.Address;
	    spotPrice: number;
	    timestamp: number;
	    transactionHash: base.Hash;
	    transactionIndex: number;
	
	    static createFrom(source: any = {}) {
	        return new Statement(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accountedFor = this.convertValues(source["accountedFor"], base.Address);
	        this.amountIn = this.convertValues(source["amountIn"], null);
	        this.amountOut = this.convertValues(source["amountOut"], null);
	        this.assetAddr = this.convertValues(source["assetAddr"], base.Address);
	        this.assetSymbol = source["assetSymbol"];
	        this.begBal = this.convertValues(source["begBal"], null);
	        this.blockNumber = source["blockNumber"];
	        this.correctingIn = this.convertValues(source["correctingIn"], null);
	        this.correctingOut = this.convertValues(source["correctingOut"], null);
	        this.correctingReason = source["correctingReason"];
	        this.decimals = source["decimals"];
	        this.endBal = this.convertValues(source["endBal"], null);
	        this.gasOut = this.convertValues(source["gasOut"], null);
	        this.internalIn = this.convertValues(source["internalIn"], null);
	        this.internalOut = this.convertValues(source["internalOut"], null);
	        this.logIndex = source["logIndex"];
	        this.minerBaseRewardIn = this.convertValues(source["minerBaseRewardIn"], null);
	        this.minerNephewRewardIn = this.convertValues(source["minerNephewRewardIn"], null);
	        this.minerTxFeeIn = this.convertValues(source["minerTxFeeIn"], null);
	        this.minerUncleRewardIn = this.convertValues(source["minerUncleRewardIn"], null);
	        this.prefundIn = this.convertValues(source["prefundIn"], null);
	        this.prevBal = this.convertValues(source["prevBal"], null);
	        this.priceSource = source["priceSource"];
	        this.recipient = this.convertValues(source["recipient"], base.Address);
	        this.selfDestructIn = this.convertValues(source["selfDestructIn"], null);
	        this.selfDestructOut = this.convertValues(source["selfDestructOut"], null);
	        this.sender = this.convertValues(source["sender"], base.Address);
	        this.spotPrice = source["spotPrice"];
	        this.timestamp = source["timestamp"];
	        this.transactionHash = this.convertValues(source["transactionHash"], base.Hash);
	        this.transactionIndex = source["transactionIndex"];
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
	export class Transaction {
	    articulatedTx?: Function;
	    blockHash: base.Hash;
	    blockNumber: number;
	    from: base.Address;
	    gas: number;
	    gasPrice: number;
	    gasUsed: number;
	    hasToken: boolean;
	    hash: base.Hash;
	    input: string;
	    isError: boolean;
	    maxFeePerGas: number;
	    maxPriorityFeePerGas: number;
	    nonce: number;
	    receipt?: Receipt;
	    timestamp: number;
	    to: base.Address;
	    traces: Trace[];
	    transactionIndex: number;
	    type: string;
	    // Go type: base
	    value: any;
	    statements?: Statement[];
	
	    static createFrom(source: any = {}) {
	        return new Transaction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.articulatedTx = this.convertValues(source["articulatedTx"], Function);
	        this.blockHash = this.convertValues(source["blockHash"], base.Hash);
	        this.blockNumber = source["blockNumber"];
	        this.from = this.convertValues(source["from"], base.Address);
	        this.gas = source["gas"];
	        this.gasPrice = source["gasPrice"];
	        this.gasUsed = source["gasUsed"];
	        this.hasToken = source["hasToken"];
	        this.hash = this.convertValues(source["hash"], base.Hash);
	        this.input = source["input"];
	        this.isError = source["isError"];
	        this.maxFeePerGas = source["maxFeePerGas"];
	        this.maxPriorityFeePerGas = source["maxPriorityFeePerGas"];
	        this.nonce = source["nonce"];
	        this.receipt = this.convertValues(source["receipt"], Receipt);
	        this.timestamp = source["timestamp"];
	        this.to = this.convertValues(source["to"], base.Address);
	        this.traces = this.convertValues(source["traces"], Trace);
	        this.transactionIndex = source["transactionIndex"];
	        this.type = source["type"];
	        this.value = this.convertValues(source["value"], null);
	        this.statements = this.convertValues(source["statements"], Statement);
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
	export class TransactionContainer {
	    items: Transaction[];
	    nItems: number;
	    address: base.Address;
	    name: string;
	    balance: string;
	    nEvents: number;
	    nTokens: number;
	    nErrors: number;
	
	    static createFrom(source: any = {}) {
	        return new TransactionContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], Transaction);
	        this.nItems = source["nItems"];
	        this.address = this.convertValues(source["address"], base.Address);
	        this.name = source["name"];
	        this.balance = source["balance"];
	        this.nEvents = source["nEvents"];
	        this.nTokens = source["nTokens"];
	        this.nErrors = source["nErrors"];
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

