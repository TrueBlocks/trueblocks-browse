export namespace app {
	
	export class AppInfo {
	    chain: string;
	    filename: string;
	    dirty: boolean;
	    meta: types.MetaData;
	    address: base.Address;
	
	    static createFrom(source: any = {}) {
	        return new AppInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.filename = source["filename"];
	        this.dirty = source["dirty"];
	        this.meta = this.convertValues(source["meta"], types.MetaData);
	        this.address = this.convertValues(source["address"], base.Address);
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
	export class ModifyData {
	    operation: string;
	    address: base.Address;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new ModifyData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.operation = source["operation"];
	        this.address = this.convertValues(source["address"], base.Address);
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

}

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

export namespace configtypes {
	
	export class ScrapeSettings {
	    appsPerChunk: number;
	    snapToGrid: number;
	    firstSnap: number;
	    unripeDist: number;
	    allowMissing?: boolean;
	    channelCount?: number;
	
	    static createFrom(source: any = {}) {
	        return new ScrapeSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.appsPerChunk = source["appsPerChunk"];
	        this.snapToGrid = source["snapToGrid"];
	        this.firstSnap = source["firstSnap"];
	        this.unripeDist = source["unripeDist"];
	        this.allowMissing = source["allowMissing"];
	        this.channelCount = source["channelCount"];
	    }
	}
	export class ChainGroup {
	    chain: string;
	    chainId: string;
	    ipfsGateway: string;
	    keyEndpoint: string;
	    localExplorer: string;
	    remoteExplorer: string;
	    rpcProvider: string;
	    symbol: string;
	    scrape: ScrapeSettings;
	
	    static createFrom(source: any = {}) {
	        return new ChainGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.chainId = source["chainId"];
	        this.ipfsGateway = source["ipfsGateway"];
	        this.keyEndpoint = source["keyEndpoint"];
	        this.localExplorer = source["localExplorer"];
	        this.remoteExplorer = source["remoteExplorer"];
	        this.rpcProvider = source["rpcProvider"];
	        this.symbol = source["symbol"];
	        this.scrape = this.convertValues(source["scrape"], ScrapeSettings);
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
	export class KeyGroup {
	    license: string;
	    apiKey: string;
	    secret: string;
	    jwt: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.license = source["license"];
	        this.apiKey = source["apiKey"];
	        this.secret = source["secret"];
	        this.jwt = source["jwt"];
	    }
	}
	export class NotifyGroup {
	    url?: string;
	    author?: string;
	
	    static createFrom(source: any = {}) {
	        return new NotifyGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.author = source["author"];
	    }
	}
	export class PinningGroup {
	    gatewayUrl: string;
	    localPinUrl: string;
	    remotePinUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new PinningGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.gatewayUrl = source["gatewayUrl"];
	        this.localPinUrl = source["localPinUrl"];
	        this.remotePinUrl = source["remotePinUrl"];
	    }
	}
	
	export class SettingsGroup {
	    cachePath: string;
	    indexPath: string;
	    defaultChain: string;
	    defaultGateway: string;
	    notify: NotifyGroup;
	
	    static createFrom(source: any = {}) {
	        return new SettingsGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.cachePath = source["cachePath"];
	        this.indexPath = source["indexPath"];
	        this.defaultChain = source["defaultChain"];
	        this.defaultGateway = source["defaultGateway"];
	        this.notify = this.convertValues(source["notify"], NotifyGroup);
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
	export class UnchainedGroup {
	    preferredPublisher: string;
	    smartContract: string;
	
	    static createFrom(source: any = {}) {
	        return new UnchainedGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.preferredPublisher = source["preferredPublisher"];
	        this.smartContract = source["smartContract"];
	    }
	}
	export class VersionGroup {
	    current: string;
	
	    static createFrom(source: any = {}) {
	        return new VersionGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current = source["current"];
	    }
	}

}

export namespace editors {
	
	export class Name {
	    address: string;
	    name: string;
	    tags: string;
	    source: string;
	    symbol: string;
	    decimals: number;
	    deleted?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Name(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = source["address"];
	        this.name = source["name"];
	        this.tags = source["tags"];
	        this.source = source["source"];
	        this.symbol = source["symbol"];
	        this.decimals = source["decimals"];
	        this.deleted = source["deleted"];
	    }
	}

}

export namespace messages {
	
	export enum Message {
	    STARTED = "Started",
	    PROGRESS = "Progress",
	    COMPLETED = "Completed",
	    CANCELED = "Canceled",
	    LOADING = "Loading",
	    LOADED = "Loaded",
	    ERROR = "Error",
	    WARNING = "Warn",
	    INFO = "Info",
	    SWITCHTAB = "SwitchTab",
	    TOGGLELAYOUT = "ToggleLayout",
	    TOGGLEACCORDION = "ToggleAccordion",
	    NAVIGATE = "Navigate",
	    REFRESH = "Refresh",
	}
	export class MessageMsg {
	    name: string;
	    address: base.Address;
	    state: string;
	    num1: number;
	    num2: number;
	    string1: string;
	    string2: string;
	    bool: boolean;
	
	    static createFrom(source: any = {}) {
	        return new MessageMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.address = this.convertValues(source["address"], base.Address);
	        this.state = source["state"];
	        this.num1 = source["num1"];
	        this.num2 = source["num2"];
	        this.string1 = source["string1"];
	        this.string2 = source["string2"];
	        this.bool = source["bool"];
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

export namespace sdk {
	
	export class SortSpec {
	    fields: string[];
	    orders: boolean[];
	
	    static createFrom(source: any = {}) {
	        return new SortSpec(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fields = source["fields"];
	        this.orders = source["orders"];
	    }
	}
	export class UpdaterItem {
	    path: string;
	    duration: number;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdaterItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.duration = source["duration"];
	        this.type = source["type"];
	    }
	}
	export class Updater {
	    name: string;
	    lastTimeStamp: number;
	    lastTotalSize: number;
	    items: UpdaterItem[];
	
	    static createFrom(source: any = {}) {
	        return new Updater(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.lastTimeStamp = source["lastTimeStamp"];
	        this.lastTotalSize = source["lastTotalSize"];
	        this.items = this.convertValues(source["items"], UpdaterItem);
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
	
	export enum DaemonState {
	    STOPPED = "Stopped",
	    RUNNING = "Running",
	    PAUSED = "Paused",
	}
	export enum WizState {
	    WELCOME = "welcome",
	    CONFIG = "config",
	    RPC = "rpc",
	    BLOOMS = "blooms",
	    INDEX = "index",
	    FINISHED = "finished",
	}
	export enum WizStep {
	    FIRST = "First",
	    PREVIOUS = "Previous",
	    NEXT = "Next",
	    FINISH = "Finish",
	}
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
	    hasConstructor: boolean;
	    hasFallback: boolean;
	    isEmpty: boolean;
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
	        this.hasConstructor = source["hasConstructor"];
	        this.hasFallback = source["hasFallback"];
	        this.isEmpty = source["isEmpty"];
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
	    hasConstructor: boolean;
	    hasFallback: boolean;
	    isEmpty: boolean;
	    isKnown: boolean;
	    lastModDate: string;
	    nEvents: number;
	    nFunctions: number;
	    name: string;
	    path: string;
	    chain: string;
	    items: Abi[];
	    largestFile: string;
	    mostEvents: string;
	    mostFunctions: string;
	    nItems: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new AbiContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.fileSize = source["fileSize"];
	        this.functions = this.convertValues(source["functions"], Function);
	        this.hasConstructor = source["hasConstructor"];
	        this.hasFallback = source["hasFallback"];
	        this.isEmpty = source["isEmpty"];
	        this.isKnown = source["isKnown"];
	        this.lastModDate = source["lastModDate"];
	        this.nEvents = source["nEvents"];
	        this.nFunctions = source["nFunctions"];
	        this.name = source["name"];
	        this.path = source["path"];
	        this.chain = source["chain"];
	        this.items = this.convertValues(source["items"], Abi);
	        this.largestFile = source["largestFile"];
	        this.mostEvents = source["mostEvents"];
	        this.mostFunctions = source["mostFunctions"];
	        this.nItems = source["nItems"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	export class RangeDates {
	    firstDate?: string;
	    firstTs?: number;
	    lastDate?: string;
	    lastTs?: number;
	
	    static createFrom(source: any = {}) {
	        return new RangeDates(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.firstDate = source["firstDate"];
	        this.firstTs = source["firstTs"];
	        this.lastDate = source["lastDate"];
	        this.lastTs = source["lastTs"];
	    }
	}
	export class ChunkRecord {
	    bloomHash: string;
	    bloomSize: number;
	    indexHash: string;
	    indexSize: number;
	    range: string;
	    rangeDates?: RangeDates;
	
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
	        this.rangeDates = this.convertValues(source["rangeDates"], RangeDates);
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
	    rangeDates?: RangeDates;
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
	        this.rangeDates = this.convertValues(source["rangeDates"], RangeDates);
	        this.ratio = source["ratio"];
	        this.recWid = source["recWid"];
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
	export class Config {
	    version: configtypes.VersionGroup;
	    settings: configtypes.SettingsGroup;
	    keys: {[key: string]: configtypes.KeyGroup};
	    pinning: configtypes.PinningGroup;
	    unchained: configtypes.UnchainedGroup;
	    chains: {[key: string]: configtypes.ChainGroup};
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = this.convertValues(source["version"], configtypes.VersionGroup);
	        this.settings = this.convertValues(source["settings"], configtypes.SettingsGroup);
	        this.keys = this.convertValues(source["keys"], configtypes.KeyGroup, true);
	        this.pinning = this.convertValues(source["pinning"], configtypes.PinningGroup);
	        this.unchained = this.convertValues(source["unchained"], configtypes.UnchainedGroup);
	        this.chains = this.convertValues(source["chains"], configtypes.ChainGroup, true);
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
	export class ConfigContainer {
	    chain: string;
	    version: configtypes.VersionGroup;
	    settings: configtypes.SettingsGroup;
	    keys: {[key: string]: configtypes.KeyGroup};
	    pinning: configtypes.PinningGroup;
	    unchained: configtypes.UnchainedGroup;
	    chains: {[key: string]: configtypes.ChainGroup};
	    items: Chain[];
	    nChains: number;
	    nItems: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new ConfigContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.version = this.convertValues(source["version"], configtypes.VersionGroup);
	        this.settings = this.convertValues(source["settings"], configtypes.SettingsGroup);
	        this.keys = this.convertValues(source["keys"], configtypes.KeyGroup, true);
	        this.pinning = this.convertValues(source["pinning"], configtypes.PinningGroup);
	        this.unchained = this.convertValues(source["unchained"], configtypes.UnchainedGroup);
	        this.chains = this.convertValues(source["chains"], configtypes.ChainGroup, true);
	        this.items = this.convertValues(source["items"], Chain);
	        this.nChains = source["nChains"];
	        this.nItems = source["nItems"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	export class Daemon {
	    name: string;
	    sleep: number;
	    color: string;
	    // Go type: time
	    started: any;
	    ticks: number;
	    state: DaemonState;
	
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
	export class DaemonIpfs {
	    name: string;
	    sleep: number;
	    color: string;
	    // Go type: time
	    started: any;
	    ticks: number;
	    state: DaemonState;
	
	    static createFrom(source: any = {}) {
	        return new DaemonIpfs(source);
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
	export class DaemonFreshen {
	    name: string;
	    sleep: number;
	    color: string;
	    // Go type: time
	    started: any;
	    ticks: number;
	    state: DaemonState;
	
	    static createFrom(source: any = {}) {
	        return new DaemonFreshen(source);
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
	export class DaemonScraper {
	    name: string;
	    sleep: number;
	    color: string;
	    // Go type: time
	    started: any;
	    ticks: number;
	    state: DaemonState;
	
	    static createFrom(source: any = {}) {
	        return new DaemonScraper(source);
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
	export class Nothing {
	    unused: string;
	
	    static createFrom(source: any = {}) {
	        return new Nothing(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.unused = source["unused"];
	    }
	}
	export class DaemonContainer {
	    chain: string;
	    name: string;
	    sleep: number;
	    color: string;
	    // Go type: time
	    started: any;
	    ticks: number;
	    state: DaemonState;
	    items: Nothing[];
	    nItems: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	    scraperController?: DaemonScraper;
	    freshenController?: DaemonFreshen;
	    ipfsController?: DaemonIpfs;
	
	    static createFrom(source: any = {}) {
	        return new DaemonContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.name = source["name"];
	        this.sleep = source["sleep"];
	        this.color = source["color"];
	        this.started = this.convertValues(source["started"], null);
	        this.ticks = source["ticks"];
	        this.state = source["state"];
	        this.items = this.convertValues(source["items"], Nothing);
	        this.nItems = source["nItems"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
	        this.scraperController = this.convertValues(source["scraperController"], DaemonScraper);
	        this.freshenController = this.convertValues(source["freshenController"], DaemonFreshen);
	        this.ipfsController = this.convertValues(source["ipfsController"], DaemonIpfs);
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
	export class Filter {
	    criteria: string;
	
	    static createFrom(source: any = {}) {
	        return new Filter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.criteria = source["criteria"];
	    }
	}
	
	export class Headers {
	    project: boolean;
	    history: boolean;
	    monitors: boolean;
	    names: boolean;
	    abis: boolean;
	    indexes: boolean;
	    manifests: boolean;
	    status: boolean;
	    settings: boolean;
	    daemons: boolean;
	    session: boolean;
	    config: boolean;
	    wizard: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Headers(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.project = source["project"];
	        this.history = source["history"];
	        this.monitors = source["monitors"];
	        this.names = source["names"];
	        this.abis = source["abis"];
	        this.indexes = source["indexes"];
	        this.manifests = source["manifests"];
	        this.status = source["status"];
	        this.settings = source["settings"];
	        this.daemons = source["daemons"];
	        this.session = source["session"];
	        this.config = source["config"];
	        this.wizard = source["wizard"];
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
	export class HistoryContainer {
	    address: base.Address;
	    balance: string;
	    chain: string;
	    items: Transaction[];
	    nErrors: number;
	    nItems: number;
	    nLogs: number;
	    nTokens: number;
	    nTotal: number;
	    name: string;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new HistoryContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.balance = source["balance"];
	        this.chain = source["chain"];
	        this.items = this.convertValues(source["items"], Transaction);
	        this.nErrors = source["nErrors"];
	        this.nItems = source["nItems"];
	        this.nLogs = source["nLogs"];
	        this.nTokens = source["nTokens"];
	        this.nTotal = source["nTotal"];
	        this.name = source["name"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	export class IndexContainer {
	    chain: string;
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
	    rangeDates?: RangeDates;
	    ratio: number;
	    recWid: number;
	    items: ChunkStats[];
	    nItems: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new IndexContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
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
	        this.rangeDates = this.convertValues(source["rangeDates"], RangeDates);
	        this.ratio = source["ratio"];
	        this.recWid = source["recWid"];
	        this.items = this.convertValues(source["items"], ChunkStats);
	        this.nItems = source["nItems"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	export class Layout {
	    header: boolean;
	    menu: boolean;
	    help: boolean;
	    footer: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Layout(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.header = source["header"];
	        this.menu = source["menu"];
	        this.help = source["help"];
	        this.footer = source["footer"];
	    }
	}
	
	export class ManifestContainer {
	    bloomsSize: number;
	    indexSize: number;
	    items: ChunkRecord[];
	    chain: string;
	    chunks: ChunkRecord[];
	    specification: string;
	    version: string;
	    nBlooms: number;
	    nIndexes: number;
	    nItems: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new ManifestContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bloomsSize = source["bloomsSize"];
	        this.indexSize = source["indexSize"];
	        this.items = this.convertValues(source["items"], ChunkRecord);
	        this.chain = source["chain"];
	        this.chunks = this.convertValues(source["chunks"], ChunkRecord);
	        this.specification = source["specification"];
	        this.version = source["version"];
	        this.nBlooms = source["nBlooms"];
	        this.nIndexes = source["nIndexes"];
	        this.nItems = source["nItems"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	
	export class Monitor {
	    address: base.Address;
	    deleted: boolean;
	    fileSize: number;
	    isEmpty: boolean;
	    isStaged: boolean;
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
	        this.isEmpty = source["isEmpty"];
	        this.isStaged = source["isStaged"];
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
	    chain: string;
	    fileSize: number;
	    items: Monitor[];
	    nDeleted: number;
	    nEmpty: number;
	    nItems: number;
	    nNamed: number;
	    nRecords: number;
	    nStaged: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new MonitorContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.fileSize = source["fileSize"];
	        this.items = this.convertValues(source["items"], Monitor);
	        this.nDeleted = source["nDeleted"];
	        this.nEmpty = source["nEmpty"];
	        this.nItems = source["nItems"];
	        this.nNamed = source["nNamed"];
	        this.nRecords = source["nRecords"];
	        this.nStaged = source["nStaged"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	    chain: string;
	    items: Name[];
	    nContracts: number;
	    nCustom: number;
	    nDeleted: number;
	    nErc20s: number;
	    nErc721s: number;
	    nItems: number;
	    nPrefund: number;
	    nRegular: number;
	    nSystem: number;
	    sizeOnDisc: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new NameContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.items = this.convertValues(source["items"], Name);
	        this.nContracts = source["nContracts"];
	        this.nCustom = source["nCustom"];
	        this.nDeleted = source["nDeleted"];
	        this.nErc20s = source["nErc20s"];
	        this.nErc721s = source["nErc721s"];
	        this.nItems = source["nItems"];
	        this.nPrefund = source["nPrefund"];
	        this.nRegular = source["nRegular"];
	        this.nSystem = source["nSystem"];
	        this.sizeOnDisc = source["sizeOnDisc"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	
	
	export class ProjectContainer {
	    chain: string;
	    historySize: number;
	    items: HistoryContainer[];
	    nAbis: number;
	    nCaches: number;
	    nIndexes: number;
	    nItems: number;
	    nManifests: number;
	    nMonitors: number;
	    nNames: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new ProjectContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.historySize = source["historySize"];
	        this.items = this.convertValues(source["items"], HistoryContainer);
	        this.nAbis = source["nAbis"];
	        this.nCaches = source["nCaches"];
	        this.nIndexes = source["nIndexes"];
	        this.nItems = source["nItems"];
	        this.nManifests = source["nManifests"];
	        this.nMonitors = source["nMonitors"];
	        this.nNames = source["nNames"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	export class Toggles {
	    layout: Layout;
	    headers: Headers;
	    daemons: Daemons;
	
	    static createFrom(source: any = {}) {
	        return new Toggles(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.layout = this.convertValues(source["layout"], Layout);
	        this.headers = this.convertValues(source["headers"], Headers);
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
	export class Window {
	    x: number;
	    y: number;
	    width: number;
	    height: number;
	    title: string;
	
	    static createFrom(source: any = {}) {
	        return new Window(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.x = source["x"];
	        this.y = source["y"];
	        this.width = source["width"];
	        this.height = source["height"];
	        this.title = source["title"];
	    }
	}
	export class Session {
	    lastChain: string;
	    lastFile: string;
	    lastFolder: string;
	    lastRoute: string;
	    lastSub: {[key: string]: string};
	    window: Window;
	    wizardStr: string;
	    toggles: Toggles;
	
	    static createFrom(source: any = {}) {
	        return new Session(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lastChain = source["lastChain"];
	        this.lastFile = source["lastFile"];
	        this.lastFolder = source["lastFolder"];
	        this.lastRoute = source["lastRoute"];
	        this.lastSub = source["lastSub"];
	        this.window = this.convertValues(source["window"], Window);
	        this.wizardStr = source["wizardStr"];
	        this.toggles = this.convertValues(source["toggles"], Toggles);
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
	export class SessionContainer {
	    chain: string;
	    items: Nothing[];
	    nItems: number;
	    lastChain: string;
	    lastFile: string;
	    lastFolder: string;
	    lastRoute: string;
	    lastSub: {[key: string]: string};
	    window: Window;
	    wizardStr: string;
	    toggles: Toggles;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new SessionContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.items = this.convertValues(source["items"], Nothing);
	        this.nItems = source["nItems"];
	        this.lastChain = source["lastChain"];
	        this.lastFile = source["lastFile"];
	        this.lastFolder = source["lastFolder"];
	        this.lastRoute = source["lastRoute"];
	        this.lastSub = source["lastSub"];
	        this.window = this.convertValues(source["window"], Window);
	        this.wizardStr = source["wizardStr"];
	        this.toggles = this.convertValues(source["toggles"], Toggles);
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	export class StatusContainer {
	    items: CacheItem[];
	    nBytes: number;
	    nFiles: number;
	    nFolders: number;
	    nItems: number;
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
	    meta?: MetaData;
	    diffs?: MetaData;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	
	    static createFrom(source: any = {}) {
	        return new StatusContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], CacheItem);
	        this.nBytes = source["nBytes"];
	        this.nFiles = source["nFiles"];
	        this.nFolders = source["nFolders"];
	        this.nItems = source["nItems"];
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
	        this.meta = this.convertValues(source["meta"], MetaData);
	        this.diffs = this.convertValues(source["diffs"], MetaData);
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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
	export class SettingsContainer {
	    chain: string;
	    items: CacheItem[];
	    nItems: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	    status: StatusContainer;
	    config: ConfigContainer;
	    session: SessionContainer;
	
	    static createFrom(source: any = {}) {
	        return new SettingsContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.items = this.convertValues(source["items"], CacheItem);
	        this.nItems = source["nItems"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
	        this.status = this.convertValues(source["status"], StatusContainer);
	        this.config = this.convertValues(source["config"], ConfigContainer);
	        this.session = this.convertValues(source["session"], SessionContainer);
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
	
	
	
	
	
	
	
	export class WizError {
	    index: number;
	    state: WizState;
	    reason: string;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new WizError(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.state = source["state"];
	        this.reason = source["reason"];
	        this.error = source["error"];
	    }
	}
	export class WizardContainer {
	    chain: string;
	    items: WizError[];
	    nItems: number;
	    updater: sdk.Updater;
	    sorts: sdk.SortSpec;
	    state: WizState;
	
	    static createFrom(source: any = {}) {
	        return new WizardContainer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.chain = source["chain"];
	        this.items = this.convertValues(source["items"], WizError);
	        this.nItems = source["nItems"];
	        this.updater = this.convertValues(source["updater"], sdk.Updater);
	        this.sorts = this.convertValues(source["sorts"], sdk.SortSpec);
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

