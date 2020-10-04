import React from 'react';

export default class Err404 extends React.Component {
    render() {
        console.log(process.env.PUBLIC_URL)
        return (
            <div>

                <div className="">
                    Page not found
                    </div>

            </div>
        );
    }
}