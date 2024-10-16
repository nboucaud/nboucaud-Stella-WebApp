// Import necessary modules
import {composeWithDevTools} from '@redux-devtools/extension';
import {
    applyMiddleware,
    legacy_createStore,
} from 'redux';
import type {
    Reducer,
    Store,
} from 'redux';
import thunk from 'redux-thunk';

import type {GlobalState} from '@mattermost/types/store';

import {createReducer} from './helpers';
import initialState from './initial_state';
import reducerRegistry from './reducer_registry';
import modalReducer from '../reducers/entities/modalReducer';
import serviceReducers from '../reducers';
import tabReducer from 'mattermost-redux/reducers/entities/tabReducer';
import urlManagerReducer from '../../../../components/browser_view/browser/browser-state'; // Import your new reducer

/**
 * Configures and constructs the redux store. Accepts the following parameters:
 * preloadedState - Any preloaded state to be applied to the store after it is initially configured.
 * appReducers - An object containing any app-specific reducer functions that the client needs.
 * getAppReducers - A function that returns the appReducer as defined above. Only used in development to enable hot reloading.
 */
export default function configureStore<S extends GlobalState>({
    appReducers,
    getAppReducers,
    preloadedState,
}: {
    appReducers: Record<string, Reducer>;
    getAppReducers: () => Record<string, Reducer>;
    preloadedState: Partial<S>;
}): Store {
    const baseState = {
        ...initialState,
        ...preloadedState,
    };

    const composeEnhancers = composeWithDevTools({
        shouldHotReload: false,
        trace: true,
        traceLimit: 25,
        autoPause: true,
    });

    const middleware = applyMiddleware(
        thunk.withExtraArgument({loaders: {}}),
    );

    const enhancers = composeEnhancers(middleware);

    // Add the urlManagerReducer to the appReducers
    const baseReducer = createReducer({
        ...serviceReducers,
        modal: modalReducer, 
        tabs: tabReducer,
        urlManager: urlManagerReducer, // Add your URL manager reducer here
        ...appReducers,
    });

    const store = legacy_createStore(
        baseReducer,
        baseState,
        enhancers
    );

    reducerRegistry.setChangeListener((reducers: Record<string, Reducer>) => {
        store.replaceReducer(createReducer(reducers, serviceReducers, appReducers));
    });

    if (module.hot) {
        // Enable Webpack hot module replacement for reducers
        module.hot.accept(() => {
            const registryReducers = reducerRegistry.getReducers();
            const nextServiceReducers = require('../reducers').default; // eslint-disable-line global-require
            const nextAppReducers = getAppReducers();

            // Ensure registryReducers comes first so that stored service/app reducers are replaced by the new ones
            store.replaceReducer(createReducer(registryReducers, nextServiceReducers, nextAppReducers));
        });
    }

    return store;
}
