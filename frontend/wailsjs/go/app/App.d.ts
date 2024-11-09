// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from '../models';
import {menu} from '../models';
import {base} from '../models';
import {app} from '../models';
import {configtypes} from '../models';
import {context} from '../models';
import {editors} from '../models';

export function AbiPage(arg1:number,arg2:number):Promise<types.AbiContainer>;

export function AbisView(arg1:menu.CallbackData):Promise<void>;

export function CancelAllContexts():Promise<void>;

export function ConfigPage(arg1:number,arg2:number):Promise<types.ConfigContainer>;

export function ConfigView(arg1:menu.CallbackData):Promise<void>;

export function DaemonPage(arg1:number,arg2:number):Promise<types.DaemonContainer>;

export function DaemonsView(arg1:menu.CallbackData):Promise<void>;

export function ExportAddress(arg1:base.Address):Promise<void>;

export function FileNew(arg1:menu.CallbackData):Promise<void>;

export function FileOpen(arg1:menu.CallbackData):Promise<void>;

export function FileSave(arg1:menu.CallbackData):Promise<void>;

export function FileSaveAs(arg1:menu.CallbackData):Promise<void>;

export function Freshen():Promise<void>;

export function GetAppInfo():Promise<app.AppInfo>;

export function GetAppTitle():Promise<string>;

export function GetChainInfo(arg1:string):Promise<types.Chain>;

export function GetChains():Promise<Array<string>>;

export function GetConfig():Promise<configtypes.Config>;

export function GetContext():Promise<context.Context>;

export function GetDaemon(arg1:string):Promise<string>;

export function GetDeferredErrors():Promise<Array<types.WizError>>;

export function GetEnv(arg1:string):Promise<string>;

export function GetExploreUrl(arg1:string,arg2:boolean,arg3:boolean):Promise<string>;

export function GetMenus():Promise<menu.Menu>;

export function GetName(arg1:base.Address):Promise<string>;

export function GetRoute():Promise<string>;

export function GetSelected():Promise<base.Address>;

export function GetSession():Promise<types.Session>;

export function GetState(arg1:string):Promise<string>;

export function GetWindow():Promise<types.Window>;

export function GoToAddress(arg1:base.Address):Promise<void>;

export function HistoryPage(arg1:number,arg2:number):Promise<types.HistoryContainer>;

export function HistoryView(arg1:menu.CallbackData):Promise<void>;

export function IndexPage(arg1:number,arg2:number):Promise<types.IndexContainer>;

export function IndexesView(arg1:menu.CallbackData):Promise<void>;

export function IsShowing(arg1:string):Promise<boolean>;

export function LoadName(arg1:string):Promise<editors.Name>;

export function Logger(arg1:Array<string>):Promise<void>;

export function ManifestPage(arg1:number,arg2:number):Promise<types.ManifestContainer>;

export function ManifestsView(arg1:menu.CallbackData):Promise<void>;

export function ModifyAbi(arg1:app.ModifyData):Promise<void>;

export function ModifyHistory(arg1:app.ModifyData):Promise<void>;

export function ModifyMonitors(arg1:app.ModifyData):Promise<void>;

export function ModifyName(arg1:app.ModifyData):Promise<void>;

export function MonitorPage(arg1:number,arg2:number):Promise<types.MonitorContainer>;

export function MonitorsView(arg1:menu.CallbackData):Promise<void>;

export function NamePage(arg1:number,arg2:number):Promise<types.NameContainer>;

export function NamesView(arg1:menu.CallbackData):Promise<void>;

export function Navigate(arg1:string,arg2:string):Promise<void>;

export function ProjectPage(arg1:number,arg2:number):Promise<types.ProjectContainer>;

export function ProjectView(arg1:menu.CallbackData):Promise<void>;

export function Reload():Promise<void>;

export function SessionPage(arg1:number,arg2:number):Promise<types.SessionContainer>;

export function SessionView(arg1:menu.CallbackData):Promise<void>;

export function SetChain(arg1:string):Promise<void>;

export function SetEnv(arg1:string,arg2:string):Promise<void>;

export function SetRoute(arg1:string,arg2:string):Promise<void>;

export function SetShowing(arg1:string,arg2:boolean):Promise<void>;

export function SettingsPage(arg1:number,arg2:number):Promise<types.SettingsContainer>;

export function SettingsView(arg1:menu.CallbackData):Promise<void>;

export function StatusPage(arg1:number,arg2:number):Promise<types.StatusContainer>;

export function StatusView(arg1:menu.CallbackData):Promise<void>;

export function StepWizard(arg1:types.WizStep):Promise<types.WizState>;

export function SystemAbout(arg1:menu.CallbackData):Promise<void>;

export function SystemQuit(arg1:menu.CallbackData):Promise<void>;

export function ToggleAccordion(arg1:menu.CallbackData):Promise<void>;

export function ToggleDaemon(arg1:string):Promise<void>;

export function ToggleFooter(arg1:menu.CallbackData):Promise<void>;

export function ToggleHeader(arg1:menu.CallbackData):Promise<void>;

export function ToggleHelp(arg1:menu.CallbackData):Promise<void>;

export function ToggleMenu(arg1:menu.CallbackData):Promise<void>;

export function ToggleNextTab(arg1:menu.CallbackData):Promise<void>;

export function TogglePrevTab(arg1:menu.CallbackData):Promise<void>;

export function WizardPage(arg1:number,arg2:number):Promise<types.WizardContainer>;

export function WizardView(arg1:menu.CallbackData):Promise<void>;
