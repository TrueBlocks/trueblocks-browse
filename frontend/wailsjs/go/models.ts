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
	    error: any;
	
	    static createFrom(source: any = {}) {
	        return new ErrorMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = this.convertValues(source["address"], base.Address);
	        this.error = source["error"];
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
	
	export enum State {
	    STOPPED = 0,
	    RUNNING = 1,
	    PAUSED = 2,
	}
	export enum Type {
	    FILESERVER = 0,
	    SCRAPER = 1,
	    MONITOR = 2,
	    API = 3,
	    IPFS = 4,
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
	
	
	export class Stats {
	    nAddresses: number;
	    nCoins: number;
	    nContracts: number;
	    nTokenSeries: number;
	    nTokenUtxo: number;
	    nTokens: number;
	    nTxns: number;
	    nUtxo: number;
	
	    static createFrom(source: any = {}) {
	        return new Stats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nAddresses = source["nAddresses"];
	        this.nCoins = source["nCoins"];
	        this.nContracts = source["nContracts"];
	        this.nTokenSeries = source["nTokenSeries"];
	        this.nTokenUtxo = source["nTokenUtxo"];
	        this.nTokens = source["nTokens"];
	        this.nTxns = source["nTxns"];
	        this.nUtxo = source["nUtxo"];
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
	    stats?: Stats;
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
	        this.stats = this.convertValues(source["stats"], Stats);
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

