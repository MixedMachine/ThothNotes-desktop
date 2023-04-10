import { Marked, Renderer } from '@ts-stack/markdown';
import highlight from 'highlight.js';

Marked.setOptions
    ({
        renderer: new Renderer,
        gfm: true,
        tables: true,
        breaks: false,
        pedantic: false,
        sanitize: false,
        smartLists: true,
        smartypants: false,
        highlight: (code, lang) => {
            return highlight.highlightAuto(code, [lang]).value;
        }
    });

export function formatContent(content: string) {
    return Marked.parse(content);
}