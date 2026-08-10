// @vitest-environment jsdom

import React from 'react';
import { render } from '@testing-library/react';
import { describe, expect, it } from 'vitest';

import RippleProvider from './RippleProvider';

describe('RippleProvider', () => {
  it('ignores document mouseleave events without throwing', () => {
    render(
      <RippleProvider>
        <button type="button" className="ripple">Import</button>
      </RippleProvider>,
    );

    expect(() => document.dispatchEvent(new MouseEvent('mouseleave'))).not.toThrow();
  });

  it('creates and clears a ripple for nested click targets', () => {
    const { getByText } = render(
      <RippleProvider>
        <button type="button" className="ripple"><span>Upload</span></button>
      </RippleProvider>,
    );
    const child = getByText('Upload');
    const button = child.closest('button');

    child.dispatchEvent(new MouseEvent('mousedown', { bubbles: true }));
    expect(button?.querySelector('.ripple-effect-span')).not.toBeNull();

    document.dispatchEvent(new MouseEvent('mouseleave'));
    expect(button?.querySelector('.ripple-effect-span')?.classList.contains('out')).toBe(true);
  });
});
