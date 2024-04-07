# MVP First Draft -> Aka was soll die Scheiße können

Functional Requirements

    1. Ein File per drag and drop verschieben zu können
        a. Möglichkeit ein File auszuwählen
        b. Möglichkeit ein File zu ziehen
        c. Möglichkeit ein File auf einen Ordner zu ziehen 
        und dann loszulassen und dann soll es im Ordner verschoben

    2. Das ganze bei 1 nur mit einem Ordner

    3. Ein File zu erstellen und einen Namen zu geben check
        a. Add File Button
        b. Eingabe Dialog für den Namen
        c. API Call lol
    
    4. Das gleiche bei 3 nur für einen Ordner

    5. Ein File in einen bestimmten Ordner erstellen CHECK
        a. Herausfinden wer gerade selectiert ist
        b. Alles bei 3 nur mit etwas anderem API Call
    
    6. PDFs hochladen check
        a. PDF per drag and drop in den Side header ablegbar
        b. oder per add button
        c. WIE ZUM HENKER läuft das mit dem Backend lol?
        d. File anzeigen

    7. 6 bitte auch für Bilder check


# create-svelte

Everything you need to build a Svelte project, powered by [`create-svelte`](https://github.com/sveltejs/kit/tree/main/packages/create-svelte).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```bash
# create a new project in the current directory
npm create svelte@latest

# create a new project in my-app
npm create svelte@latest my-app
```

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://kit.svelte.dev/docs/adapters) for your target environment.
