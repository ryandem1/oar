import type {Outcome, Analysis, Resolution} from "./consts";

export interface Test {
    id: bigint
    summary: string
    outcome: Outcome
    analysis: Analysis
    resolution: Resolution
    doc: Record<string, any>
}
