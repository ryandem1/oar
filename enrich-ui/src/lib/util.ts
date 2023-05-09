export const toTitleCase = (str: string): string => {
	return str.toLowerCase().replace(/(?:^|\s)\w/g, (match) => match.toUpperCase());
};
